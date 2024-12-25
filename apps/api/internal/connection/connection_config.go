package connection

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionManager struct {
	connections map[string]interface{}
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]interface{}),
	}
}

func (cm *ConnectionManager) Connect(config ConnectionConfig) error {
	switch config.Type {
	case "mysql":
		return cm.connectMySQL(config)
	case "postgresql":
		return cm.connectPostgres(config)
	case "mongodb":
		return cm.connectMongoDB(config)
	default:
		return fmt.Errorf("unsupported database type: %s", config.Type)
	}
}

func (cm *ConnectionManager) connectMySQL(config ConnectionConfig) error {
	sslMode := "false"
	if config.SSL {
		sslMode = "true"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=%s",
		config.Username, config.Password, config.Host, config.Port, config.Database, sslMode)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	cm.connections[config.ID] = db
	return nil
}

func (cm *ConnectionManager) connectPostgres(config ConnectionConfig) error {
	sslMode := "disable"
	if config.SSL {
		sslMode = "require"
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.Database, sslMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	cm.connections[config.ID] = db
	return nil
}

func (cm *ConnectionManager) connectMongoDB(config ConnectionConfig) error {
	ctx := context.Background()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	cm.connections[config.ID] = client
	return nil
}

func (cm *ConnectionManager) Disconnect(id string) error {
	conn, exists := cm.connections[id]
	if !exists {
		return fmt.Errorf("connection not found: %s", id)
	}

	switch c := conn.(type) {
	case *sql.DB:
		return c.Close()
	case *mongo.Client:
		return c.Disconnect(context.Background())
	default:
		return fmt.Errorf("unknown connection type for id: %s", id)
	}
}
