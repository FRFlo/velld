package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dendianugerah/velld/internal/auth"
	"github.com/dendianugerah/velld/internal/backup"
	"github.com/dendianugerah/velld/internal/common"
	"github.com/dendianugerah/velld/internal/connection"
	"github.com/dendianugerah/velld/internal/database"
	"github.com/dendianugerah/velld/internal/middleware"
	"github.com/dendianugerah/velld/internal/notification"
	"github.com/dendianugerah/velld/internal/settings"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	if jwtSecret == "" || encryptionKey == "" {
		log.Fatal("JWT_SECRET and ENCRYPTION_KEY must be set in .env file")
	}

	dbPath := filepath.Join("internal", "database", "velld.db")
	db, err := database.Init(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	connManager := connection.NewConnectionManager()

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, jwtSecret)

	cryptoService, err := common.NewEncryptionService(encryptionKey)
	if err != nil {
		log.Fatal(err)
	}

	connRepo := connection.NewConnectionRepository(db, cryptoService)
	connService := connection.NewConnectionService(connRepo, connManager)

	connHandler := connection.NewConnectionHandler(connService)
	authHandler := auth.NewAuthHandler(authService)

	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// Public routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST", "OPTIONS")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(authMiddleware.RequireAuth)
	protected.HandleFunc("/auth/profile", authHandler.GetProfile).Methods("GET", "OPTIONS")

	protected.HandleFunc("/connections/test", connHandler.TestConnection).Methods("POST", "OPTIONS")
	protected.HandleFunc("/connections", connHandler.SaveConnection).Methods("POST", "OPTIONS")
	protected.HandleFunc("/connections", connHandler.ListConnections).Methods("GET", "OPTIONS")
	protected.HandleFunc("/connections", connHandler.UpdateConnection).Methods("PUT", "OPTIONS")

	// Configure backup tool paths from environment
	toolPaths := map[string]string{
		"postgresql": os.Getenv("POSTGRESQL_BIN_PATH"),
		"mysql":      os.Getenv("MYSQL_BIN_PATH"),
		"mariadb":    os.Getenv("MARIADB_BIN_PATH"),
		"mongodb":    os.Getenv("MONGODB_BIN_PATH"),
	}

	backupRepo := backup.NewBackupRepository(db)
	settingsRepo := settings.NewSettingsRepository(db)
	notificationRepo := notification.NewNotificationRepository(db)

	backupService := backup.NewBackupService(
		connRepo,
		"./backups",
		toolPaths,
		backupRepo,
		settingsRepo,
		notificationRepo,
		cryptoService,
	)

	backupHandler := backup.NewBackupHandler(backupService)

	protected.HandleFunc("/backups/stats", backupHandler.GetBackupStats).Methods("GET", "OPTIONS")
	protected.HandleFunc("/backups/schedule", backupHandler.ScheduleBackup).Methods("POST", "OPTIONS")
	protected.HandleFunc("/backups", backupHandler.CreateBackup).Methods("POST", "OPTIONS")
	protected.HandleFunc("/backups", backupHandler.ListBackups).Methods("GET", "OPTIONS")
	protected.HandleFunc("/backups/{id}", backupHandler.GetBackup).Methods("GET", "OPTIONS")
	protected.HandleFunc("/backups/{id}/download", backupHandler.DownloadBackup).Methods("GET", "OPTIONS")
	protected.HandleFunc("/backups/{connection_id}/schedule/disable", backupHandler.DisableBackupSchedule).Methods("POST", "OPTIONS")
	protected.HandleFunc("/backups/{connection_id}/schedule", backupHandler.UpdateBackupSchedule).Methods("PUT", "OPTIONS")

	settingsService := settings.NewSettingsService(settingsRepo, cryptoService)
	settingsHandler := settings.NewSettingsHandler(settingsService)

	protected.HandleFunc("/settings", settingsHandler.GetSettings).Methods("GET", "OPTIONS")
	protected.HandleFunc("/settings", settingsHandler.UpdateSettings).Methods("PUT", "OPTIONS")

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
