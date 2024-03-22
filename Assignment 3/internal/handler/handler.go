package handler

import (
	"encoding/json"
	"net/http"

	"assignment-3/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := h.service.GenerateRandomStatus()
	statusText := h.service.DetermineStatus(status)

	err := h.service.UpdateJSONStatus(status, statusText)
	if err != nil {
		http.Error(w, "Gagal memperbarui status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  status,
		"message": statusText,
	})
}
