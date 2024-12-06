package handler

import (
	"encoding/json"
	"encrypt-decrypt/internal/entity"
	"encrypt-decrypt/internal/service"
	"net/http"
)

type DecryptHandler struct {
	service service.DecryptService
}

func NewDecryptHandler(service service.DecryptService) *DecryptHandler {
	return &DecryptHandler{service: service}
}

func (h *DecryptHandler) DecryptData(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.Request
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := h.service.Decrypt(requestBody.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(entity.NewResponse(*result, requestBody.Text))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
