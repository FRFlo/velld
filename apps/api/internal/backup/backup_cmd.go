package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/dendianugerah/velld/internal/connection"
)

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
