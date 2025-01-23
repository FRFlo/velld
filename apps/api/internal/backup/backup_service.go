package backup

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dendianugerah/velld/internal/connection"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

type BackupService struct {
	connStorage *connection.ConnectionRepository
	backupDir   string
	toolPaths   map[string]string
	backupRepo  *BackupRepository
	cronManager *cron.Cron
	cronEntries map[string]cron.EntryID // map[scheduleID]entryID
}

func NewBackupService(cs *connection.ConnectionRepository, backupDir string, toolPaths map[string]string, backupRepo *BackupRepository) *BackupService {
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		panic(err)
	}

	// Initialize default paths if not provided
	if toolPaths == nil {
		toolPaths = make(map[string]string)
	}

	cronManager := cron.New(cron.WithSeconds())
	service := &BackupService{
		connStorage: cs,
		backupDir:   backupDir,
		toolPaths:   toolPaths,
		backupRepo:  backupRepo,
		cronManager: cronManager,
		cronEntries: make(map[string]cron.EntryID),
	}

	// Recover existing schedules before starting the cron manager
	if err := service.recoverSchedules(); err != nil {
		fmt.Printf("Error recovering schedules: %v\n", err)
	}

	cronManager.Start()
	return service
}

func (s *BackupService) recoverSchedules() error {
	schedules, err := s.backupRepo.GetAllActiveSchedules()
	if err != nil {
		return fmt.Errorf("failed to get active schedules: %v", err)
	}

	now := time.Now()
	for _, schedule := range schedules {
		scheduleID := schedule.ID.String()

		// Check if we missed any backups
		if schedule.NextRunTime != nil && schedule.NextRunTime.Before(now) {
			// Execute a backup immediately for missed schedule
			go s.executeCronBackup(schedule)
		}

		// Re-register the cron job
		entryID, err := s.cronManager.AddFunc(schedule.CronSchedule, func() {
			s.executeCronBackup(schedule)
		})
		if err != nil {
			fmt.Printf("Error re-registering schedule %s: %v\n", scheduleID, err)
			continue
		}

		s.cronEntries[scheduleID] = entryID
	}

	return nil
}

func (s *BackupService) CreateBackup(connectionID string) (*Backup, error) {
	conn, err := s.connStorage.GetConnection(connectionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection: %v", err)
	}

	if err := s.verifyBackupTools(conn.Type); err != nil {
		return nil, err
	}

	backupID := uuid.New()
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s.sql", conn.DatabaseName, timestamp)
	backupPath := filepath.Join(s.backupDir, filename)

	backup := &Backup{
		ID:           backupID,
		ConnectionID: connectionID,
		StartedTime:  time.Now(),
		Status:       "in_progress",
		Path:         backupPath,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	var cmd *exec.Cmd
	switch conn.Type {
	case "postgresql":
		cmd = s.createPgDumpCmd(conn, backupPath)
	case "mysql", "mariadb":
		cmd = s.createMySQLDumpCmd(conn, backupPath)
	case "mongodb":
		cmd = s.createMongoDumpCmd(conn, backupPath)
	default:
		return nil, fmt.Errorf("unsupported database type for backup: %s", conn.Type)
	}

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("backup failed: %v", err)
	}

	// Get file size
	fileInfo, err := os.Stat(backupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get backup file info: %v", err)
	}

	backup.Size = fileInfo.Size()
	backup.Status = "completed"
	now := time.Now()
	backup.CompletedTime = &now

	if err := s.backupRepo.CreateBackup(backup); err != nil {
		return nil, fmt.Errorf("failed to save backup: %v", err)
	}

	return backup, nil
}

func (s *BackupService) GetBackup(id string) (*Backup, error) {
	return s.backupRepo.GetBackup(id)
}

func (s *BackupService) GetAllBackupsWithPagination(opts BackupListOptions) ([]*BackupList, int, error) {
	if opts.Limit <= 0 {
		opts.Limit = 10
	}
	if opts.Limit > 100 {
		opts.Limit = 100
	}
	if opts.Offset < 0 {
		opts.Offset = 0
	}

	return s.backupRepo.GetAllBackupsWithPagination(opts)
}

func (s *BackupService) verifyBackupTools(dbType string) error {
	requiredTools := map[string][]string{
		"postgresql": {"pg_dump"},
		"mysql":      {"mysqldump"},
		"mariadb":    {"mysqldump"},
		"mongodb":    {"mongodump"},
	}

	tools, exists := requiredTools[dbType]
	if !exists {
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	for _, tool := range tools {
		if err := s.verifyTool(dbType, tool); err != nil {
			return err
		}
	}
	return nil
}

func (s *BackupService) verifyTool(dbType, toolName string) error {
	if customPath, exists := s.toolPaths[dbType]; exists {
		toolPath := filepath.Join(customPath, getPlatformExecutableName(toolName))
		if _, err := os.Stat(toolPath); err == nil {
			return nil
		}
	}

	if path, err := exec.LookPath(getPlatformExecutableName(toolName)); err == nil {
		s.toolPaths[dbType] = filepath.Dir(path)
		return nil
	}

	return fmt.Errorf("required tool %s is not installed or not in configured paths", toolName)
}

func getPlatformExecutableName(base string) string {
	if runtime.GOOS == "windows" {
		return base + ".exe"
	}
	return base
}

func (s *BackupService) createPgDumpCmd(conn *connection.StoredConnection, outputPath string) *exec.Cmd {
	pgDumpPath := filepath.Join(s.toolPaths["postgresql"], getPlatformExecutableName("pg_dump"))
	cmd := exec.Command(pgDumpPath,
		"-h", conn.Host,
		"-p", fmt.Sprintf("%d", conn.Port),
		"-U", conn.Username,
		"-d", conn.DatabaseName,
		"-f", outputPath,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", conn.Password))
	return cmd
}

func (s *BackupService) createMySQLDumpCmd(conn *connection.StoredConnection, outputPath string) *exec.Cmd {
	mysqlDumpPath := filepath.Join(s.toolPaths[conn.Type], getPlatformExecutableName("mysqldump"))
	cmd := exec.Command(mysqlDumpPath,
		"-h", conn.Host,
		"-P", fmt.Sprintf("%d", conn.Port),
		"-u", conn.Username,
		fmt.Sprintf("-p%s", conn.Password),
		conn.DatabaseName,
		"-r", outputPath,
	)
	return cmd
}

func (s *BackupService) createMongoDumpCmd(conn *connection.StoredConnection, outputPath string) *exec.Cmd {
	mongoDumpPath := filepath.Join(s.toolPaths["mongodb"], getPlatformExecutableName("mongodump"))

	args := []string{
		"--host", conn.Host,
		"--port", fmt.Sprintf("%d", conn.Port),
		"--db", conn.DatabaseName,
		"--out", filepath.Dir(outputPath),
	}

	if conn.Username != "" {
		args = append(args, "--username", conn.Username)
	}

	if conn.Password != "" {
		args = append(args, "--password", conn.Password)
	}

	cmd := exec.Command(mongoDumpPath, args...)
	return cmd
}

func (s *BackupService) ScheduleBackup(req *ScheduleBackupRequest) error {
	// First check if a schedule already exists for this connection
	existingSchedule, err := s.backupRepo.GetBackupSchedule(req.ConnectionID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check existing schedule: %v", err)
	}

	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := parser.Parse(req.CronSchedule)
	if err != nil {
		return fmt.Errorf("invalid cron schedule: %v", err)
	}

	nextRun := schedule.Next(time.Now())

	if existingSchedule != nil {
		// Update existing schedule
		existingSchedule.Enabled = true
		existingSchedule.CronSchedule = req.CronSchedule
		existingSchedule.RetentionDays = req.RetentionDays
		existingSchedule.NextRunTime = &nextRun
		existingSchedule.UpdatedAt = time.Now()

		if err := s.backupRepo.UpdateBackupSchedule(existingSchedule); err != nil {
			return fmt.Errorf("failed to update backup schedule: %v", err)
		}

		// Update cron job
		scheduleID := existingSchedule.ID.String()
		if oldEntryID, exists := s.cronEntries[scheduleID]; exists {
			s.cronManager.Remove(oldEntryID)
		}

		entryID, err := s.cronManager.AddFunc(req.CronSchedule, func() {
			s.executeCronBackup(existingSchedule)
		})
		if err != nil {
			return fmt.Errorf("failed to schedule backup: %v", err)
		}

		s.cronEntries[scheduleID] = entryID
		return nil
	}

	// Create new schedule if none exists
	backupSchedule := &BackupSchedule{
		ID:            uuid.New(),
		ConnectionID:  req.ConnectionID,
		Enabled:       true,
		CronSchedule:  req.CronSchedule,
		RetentionDays: req.RetentionDays,
		NextRunTime:   &nextRun,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := s.backupRepo.CreateBackupSchedule(backupSchedule); err != nil {
		return fmt.Errorf("failed to save backup schedule: %v", err)
	}

	scheduleID := backupSchedule.ID.String()
	entryID, err := s.cronManager.AddFunc(req.CronSchedule, func() {
		s.executeCronBackup(backupSchedule)
	})
	if err != nil {
		return fmt.Errorf("failed to schedule backup: %v", err)
	}

	s.cronEntries[scheduleID] = entryID
	return nil
}

func (s *BackupService) executeCronBackup(schedule *BackupSchedule) {
	backup, err := s.CreateBackup(schedule.ConnectionID)
	if err != nil {
		// Log error but continue with schedule update
		fmt.Printf("Error executing scheduled backup: %v\n", err)
	} else {
		// Update backup with schedule ID and status
		scheduleIDStr := schedule.ID.String()
		if err := s.backupRepo.UpdateBackupStatusAndSchedule(backup.ID.String(), backup.Status, scheduleIDStr); err != nil {
			fmt.Printf("Error updating backup status and schedule: %v\n", err)
		}
	}

	// Update schedule's next run time and last backup time
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	cronSchedule, _ := parser.Parse(schedule.CronSchedule)
	nextRun := cronSchedule.Next(time.Now())
	schedule.NextRunTime = &nextRun
	now := time.Now()
	schedule.LastBackupTime = &now
	schedule.UpdatedAt = now

	if err := s.backupRepo.UpdateBackupSchedule(schedule); err != nil {
		fmt.Printf("Error updating backup schedule: %v\n", err)
	}

	if schedule.RetentionDays > 0 {
		s.cleanupOldBackups(schedule.ConnectionID, schedule.RetentionDays)
	}
}

func (s *BackupService) cleanupOldBackups(connectionID string, retentionDays int) {
	cutoffTime := time.Now().AddDate(0, 0, -retentionDays)
	oldBackups, err := s.backupRepo.GetBackupsOlderThan(connectionID, cutoffTime)
	if err != nil {
		return
	}

	for _, backup := range oldBackups {
		os.Remove(backup.Path)
		s.backupRepo.DeleteBackup(backup.ID.String())
	}
}

func (s *BackupService) DisableBackupSchedule(connectionID string) error {
	schedule, err := s.backupRepo.GetBackupSchedule(connectionID)
	if err != nil {
		return err
	}

	scheduleID := schedule.ID.String()
	if entryID, exists := s.cronEntries[scheduleID]; exists {
		s.cronManager.Remove(entryID)
		delete(s.cronEntries, scheduleID)
	}

	schedule.Enabled = false
	schedule.UpdatedAt = time.Now()
	if err := s.backupRepo.UpdateBackupSchedule(schedule); err != nil {
		return err
	}

	return nil
}

func (s *BackupService) UpdateBackupSchedule(connectionID string, req *UpdateScheduleRequest) error {
	schedule, err := s.backupRepo.GetBackupSchedule(connectionID)
	if err != nil {
		return err
	}

	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	_, err = parser.Parse(req.CronSchedule)
	if err != nil {
		return fmt.Errorf("invalid cron schedule: %v", err)
	}

	schedule.CronSchedule = req.CronSchedule
	schedule.RetentionDays = req.RetentionDays
	err = s.backupRepo.UpdateBackupSchedule(schedule)
	if err != nil {
		return err
	}

	// Remove old cron job
	if entryID, ok := s.cronEntries[schedule.ID.String()]; ok {
		s.cronManager.Remove(entryID)
		delete(s.cronEntries, schedule.ID.String())
	}

	// Add new cron job
	entryID, err := s.cronManager.AddFunc(schedule.CronSchedule, func() {
		s.executeCronBackup(schedule)
	})
	if err != nil {
		return fmt.Errorf("failed to register cron job: %v", err)
	}

	// Store the new entry ID
	s.cronEntries[schedule.ID.String()] = entryID

	return nil
}
