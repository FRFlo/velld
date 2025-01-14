package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dendianugerah/velld/internal/connection"
	"github.com/google/uuid"
)

type BackupService struct {
	connStorage *connection.ConnectionStorage
	backupDir   string
	toolPaths   map[string]string
	backupRepo  *BackupRepository
}

func NewBackupService(cs *connection.ConnectionStorage, backupDir string, toolPaths map[string]string, backupRepo *BackupRepository) *BackupService {
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		panic(err)
	}

	// Initialize default paths if not provided
	if toolPaths == nil {
		toolPaths = make(map[string]string)
	}

	return &BackupService{
		connStorage: cs,
		backupDir:   backupDir,
		toolPaths:   toolPaths,
		backupRepo:  backupRepo,
	}
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

func (s *BackupService) GetAllBackups(userID uuid.UUID) ([]*BackupList, error) {
	return s.backupRepo.GetAllBackups(userID)
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

	// if conn.AuthenticationDatabase != "" {
	// 	args = append(args, "--authenticationDatabase", conn.AuthenticationDatabase)
	// }

	cmd := exec.Command(mongoDumpPath, args...)
	return cmd
}

func (s *BackupService) ScheduleBackup(backup *Backup) error {

	// TODO: Implement cron scheduling logic here
	// This would typically involve:
	// 1. Validating the cron expression
	// 2. Creating a cron job
	// 3. Storing the schedule in the database
	// 4. Calculating and setting the NextRunTime

	return nil
}
