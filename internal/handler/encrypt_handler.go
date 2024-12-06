package handler

import (
	"encoding/json"
	"encrypt-decrypt/internal/entity"
	"encrypt-decrypt/internal/service"
	"net/http"
)

type EncryptHandler struct {
	service service.EncryptService
}

func NewEncryptHandler(service service.EncryptService) *EncryptHandler {
	return &EncryptHandler{service: service}
}

func (h *EncryptHandler) EncryptData(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.Request
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := h.service.Encrypt(requestBody.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(entity.NewResponse(requestBody.Text, *result))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
