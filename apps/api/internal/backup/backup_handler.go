package backup

import (
	"encoding/json"
	"net/http"
)

type BackupHandler struct {
	backupService *BackupService
}

func NewBackupHandler(bs *BackupService) *BackupHandler {
	return &BackupHandler{
		backupService: bs,
	}
}

func (h *BackupHandler) CreateBackup(w http.ResponseWriter, r *http.Request) {
	var req BackupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	backup, err := h.backupService.CreateBackup(req.ConnectionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(backup)
}
