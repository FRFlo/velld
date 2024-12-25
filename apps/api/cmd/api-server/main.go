package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dendianugerah/velld/internal/auth"
	"github.com/dendianugerah/velld/internal/common"
	"github.com/dendianugerah/velld/internal/database"
	"github.com/dendianugerah/velld/internal/connection"
	"github.com/dendianugerah/velld/internal/middleware"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
)

const (
	migrationsDir = "./internal/database/migrations"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := runMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	} else {
		log.Println("Migrations applied successfully")
	}

	connManager := database.NewConnectionManager()

	// Initialize repositories
	authRepo := auth.NewAuthRepository(db)

	// Initialize services with repositories
	authService := auth.NewAuthService(authRepo, jwtSecret)

	cryptoService, err := common.NewEncryptionService(encryptionKey)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize storage with encryption
	connStorage := connection.NewConnectionStorage(db, cryptoService)

	// Initialize handlers
	connHandler := connection.NewConnectionHandler(connManager, connStorage)
	authHandler := auth.NewAuthHandler(authService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	// Setup router
	r := mux.NewRouter()

	// CORS middleware
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
	protected.HandleFunc("/connections/test", connHandler.TestConnection).Methods("POST", "OPTIONS")
	protected.HandleFunc("/connections", connHandler.SaveConnection).Methods("POST", "OPTIONS")
	protected.HandleFunc("/connections", connHandler.ListConnections).Methods("GET", "OPTIONS")

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func runMigrations(db *sql.DB) error {
	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}
	return nil
}
