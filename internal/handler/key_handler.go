package handler

import (
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
	if err := h.service.CreateNewKey(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *KeyHandler) GetKey(w http.ResponseWriter, r *http.Request) {
	key, err := h.service.GetKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(*key))
}
