package connection

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type ConnectionHandler struct {
	connManager *ConnectionManager
	connStorage *ConnectionStorage
}

func NewConnectionHandler(cm *ConnectionManager, cs *ConnectionStorage) *ConnectionHandler {
	return &ConnectionHandler{
		connManager: cm,
		connStorage: cs,
	}
}

func (h *ConnectionHandler) TestConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.connManager.Connect(config)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"isConnected": false,
			"error":       err.Error(),
		})
		return
	}

	defer h.connManager.Disconnect(config.ID)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"isConnected": true,
		"lastSync":    "",
	})
}

func (h *ConnectionHandler) SaveConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if config.ID == "" {
		config.ID = uuid.New().String()
	}

	if err := h.connManager.Connect(config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer h.connManager.Disconnect(config.ID)

	dbSize, err := h.connManager.GetDatabaseSize(config.ID)
	if err != nil {
		fmt.Printf("Failed to get database size: %v\n", err)
	}

	userClaims, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		http.Error(w, "invalid user claims", http.StatusBadRequest)
		return
	}

	userIDStr := fmt.Sprintf("%v", userClaims["user_id"])
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storedConn := StoredConnection{
		ID:           config.ID,
		Name:         config.Name,
		Type:         config.Type,
		Host:         config.Host,
		Port:         config.Port,
		Username:     config.Username,
		Password:     config.Password,
		DatabaseName: config.Database,
		SSL:          config.SSL,
		UserID:       userID,
		Status:       "connected",
		DatabaseSize: dbSize,
	}

	if err := h.connStorage.SaveConnection(storedConn); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(storedConn)
}

func (h *ConnectionHandler) ListConnections(w http.ResponseWriter, r *http.Request) {
	userClaims, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		http.Error(w, "invalid user claims", http.StatusBadRequest)
		return
	}

	userIDStr := fmt.Sprintf("%v", userClaims["user_id"])
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connections, err := h.connStorage.ListConnections(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(connections)
}

func (h *ConnectionHandler) GetConnectionStats(w http.ResponseWriter, r *http.Request) {
    userClaims, ok := r.Context().Value("user").(jwt.MapClaims)
    if !ok {
        http.Error(w, "invalid user claims", http.StatusBadRequest)
        return
    }

    userIDStr := fmt.Sprintf("%v", userClaims["user_id"])
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    stats, err := h.connStorage.GetConnectionStats(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}
