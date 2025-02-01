package connection

import (
	"encoding/json"
	"net/http"

	"github.com/dendianugerah/velld/internal/common"
	"github.com/dendianugerah/velld/internal/common/response"
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
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := common.GetUserIDFromContext(r.Context())
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	storedConn, err := h.service.SaveConnection(config, userID)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(storedConn)
}

func (h *ConnectionHandler) ListConnections(w http.ResponseWriter, r *http.Request) {
	userID, err := common.GetUserIDFromContext(r.Context())
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	connections, err := h.service.ListConnections(userID)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(connections)
}

func (h *ConnectionHandler) UpdateConnection(w http.ResponseWriter, r *http.Request) {
	var config ConnectionConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := common.GetUserIDFromContext(r.Context())
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	storedConn, err := h.service.UpdateConnection(config, userID)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storedConn)
}
