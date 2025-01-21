package connection

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type ConnectionHandler struct {
	service *ConnectionService
}

func NewConnectionHandler(service *ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		service: service,
	}
}

func (h *ConnectionHandler) TestConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isConnected, err := h.service.TestConnection(config)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"isConnected": false,
			"error":       err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"isConnected": isConnected,
		"lastSync":    "",
	})
}

func (h *ConnectionHandler) SaveConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	storedConn, err := h.service.SaveConnection(config, userID)
	if err != nil {
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

	connections, err := h.service.ListConnections(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
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

	stats, err := h.service.GetConnectionStats(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (h *ConnectionHandler) UpdateConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	storedConn, err := h.service.UpdateConnection(config, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storedConn)
}
