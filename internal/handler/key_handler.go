package handler

import (
	"encoding/json"
	"encrypt-decrypt/internal/entity"
	"encrypt-decrypt/internal/service"
	"net/http"
)

type KeyHandler struct {
	service service.KeyService
}

func NewKeyHandler(service service.KeyService) *KeyHandler {
	return &KeyHandler{service: service}
}

func (h *KeyHandler) CreateNewKey(w http.ResponseWriter, r *http.Request) {
	key, err := h.service.CreateNewKey()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(entity.NewKeyResponse(key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *KeyHandler) GetKey(w http.ResponseWriter, r *http.Request) {
	key, err := h.service.GetKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(entity.NewKeyResponse(key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
