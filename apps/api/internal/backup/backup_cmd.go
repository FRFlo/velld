package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dendianugerah/velld/internal/common"
	"github.com/dendianugerah/velld/internal/connection"
	"github.com/google/uuid"
)

var requiredTools = map[string]string{
	"postgresql": "pg_dump",
	"mysql":      "mysqldump",
	"mariadb":    "mysqldump",
	"mongodb":    "mongodump",
}

func (s *BackupService) verifyBackupTools(dbType string) error {
	if _, exists := requiredTools[dbType]; !exists {
		return fmt.Errorf("unsupported database type: %s", dbType)
	}
	return nil
}

func (s *BackupService) findDatabaseBinaryPath(dbType string, userID uuid.UUID) string {
	userPath, err := s.settingsRepo.GetDatabaseBinaryPath(dbType, userID)
	if err != nil || userPath == "" {
		return common.FindBinaryPath(dbType, requiredTools[dbType], nil)
	}

	if path := common.FindBinaryPath(dbType, requiredTools[dbType], &userPath); path != "" {
		if path != userPath {
			fmt.Printf("Warning: User-specified path %s failed, using found path: %s\n", userPath, path)
		}
		return path
	}

	return ""
}

func (s *BackupService) createPgDumpCmd(conn *connection.StoredConnection, outputPath string, userID uuid.UUID) *exec.Cmd {
	binaryPath := s.findDatabaseBinaryPath("postgresql", userID)
	if binaryPath == "" {
		return nil
	}

	binPath := filepath.Join(binaryPath, common.GetPlatformExecutableName(requiredTools["postgresql"]))

	cmd := exec.Command(binPath,
		"-h", conn.Host,
		"-p", fmt.Sprintf("%d", conn.Port),
		"-U", conn.Username,
		"-d", conn.DatabaseName,
		"-f", outputPath,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", conn.Password))
	return cmd
}

func (s *BackupService) createMySQLDumpCmd(conn *connection.StoredConnection, outputPath string, userID uuid.UUID) *exec.Cmd {
	binaryPath := s.findDatabaseBinaryPath(conn.Type, userID)
	if binaryPath == "" {
		return nil
	}

	binPath := filepath.Join(binaryPath, common.GetPlatformExecutableName(requiredTools[conn.Type]))
	cmd := exec.Command(binPath,
		"-h", conn.Host,
		"-P", fmt.Sprintf("%d", conn.Port),
		"-u", conn.Username,
		fmt.Sprintf("-p%s", conn.Password),
		conn.DatabaseName,
		"-r", outputPath,
	)
	return cmd
}

func (s *BackupService) createMongoDumpCmd(conn *connection.StoredConnection, outputPath string, userID uuid.UUID) *exec.Cmd {
	binaryPath := s.findDatabaseBinaryPath("mongodb", userID)
	if binaryPath == "" {
		return nil
	}

	binPath := filepath.Join(binaryPath, common.GetPlatformExecutableName(requiredTools["mongodb"]))
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

	return exec.Command(binPath, args...)
}
