package connection

import (
	"encoding/json"
	"net/http"

	"github.com/dendianugerah/velld/internal/database"
	"github.com/google/uuid"
)

type ConnectionHandler struct {
	connManager *database.ConnectionManager
	connStorage *ConnectionStorage
}

func NewConnectionHandler(cm *database.ConnectionManager, cs *ConnectionStorage) *ConnectionHandler {
	return &ConnectionHandler{
		connManager: cm,
		connStorage: cs,
	}
}

func (h *ConnectionHandler) TestConnection(w http.ResponseWriter, r *http.Request) {
	var config database.ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.connManager.Connect(config)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"isConnected": false,
			"error":      err.Error(),
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
	var config database.ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate UUID if ID is not provided
	if config.ID == "" {
		config.ID = uuid.New().String()
	}

	// Test the connection first
	if err := h.connManager.Connect(config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer h.connManager.Disconnect(config.ID)

	// Get user ID from context
	userClaims := r.Context().Value("user").(map[string]interface{})
	userID := int(userClaims["user_id"].(float64))

	// Save to database
	storedConn := StoredConnection{
		ID:           config.ID,
		Name:         config.Name,
		Type:         config.Type,
		Host:         config.Host,
		Port:         config.Port,
		Username:     config.Username,
		Password:     config.Password,
		DatabaseName: config.Database,
		SSL:         config.SSL,
		UserID:      userID,
		Status:      "active",
	}

	if err := h.connStorage.SaveConnection(storedConn); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(storedConn)
}

func (h *ConnectionHandler) ListConnections(w http.ResponseWriter, r *http.Request) {
	userClaims := r.Context().Value("user").(map[string]interface{})
	userID := int(userClaims["user_id"].(float64))

	connections, err := h.connStorage.ListConnections(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(connections)
}
