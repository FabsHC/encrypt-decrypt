package service

import (
	"crypto/rand"
	"encoding/hex"
	"encrypt-decrypt/internal/repository"
)

type (
	KeyService interface {
		CreateNewKey() error
		GetKey() (*string, error)
	}

	keyService struct {
		repo repository.KeyRepository
	}
)

func NewKeyService(repo repository.KeyRepository) KeyService {
	return &keyService{
		repo: repo,
	}
}

func (s *keyService) CreateNewKey() error {
	key := make([]byte, 32) // 32 bytes para HMAC-SHA256
	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	return s.repo.Create(hex.EncodeToString(key))
}

func (s *keyService) GetKey() (*string, error) {
	return s.repo.Get()
}
