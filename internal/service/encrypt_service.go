package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encrypt-decrypt/internal/repository"
	"io"
)

type (
	EncryptService interface {
		Encrypt(text string) (*string, error)
	}

	encryptService struct {
		repo repository.KeyRepository
	}
)

func NewEncryptService(repo repository.KeyRepository) EncryptService {
	return &encryptService{
		repo: repo,
	}
}

func (e *encryptService) Encrypt(text string) (*string, error) {
	// Obtém a chave do repositório
	key, err := e.repo.Get()
	if err != nil {
		return nil, err
	}

	// Decodifica a chave (em formato hexadecimal) para bytes
	bKey, err := hex.DecodeString(*key)
	if err != nil {
		return nil, err
	}

	// Cria o cifrador AES
	c, err := aes.NewCipher(bKey)
	if err != nil {
		return nil, err
	}

	// Configura o GCM (Galois/Counter Mode)
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	// Gera um nonce aleatório
	bNonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, bNonce); err != nil {
		return nil, err
	}

	// Criptografa o texto diretamente (sem hexificar)
	bData := gcm.Seal(bNonce, bNonce, []byte(text), nil)

	// Codifica o resultado criptografado para Base64
	strData := base64.StdEncoding.EncodeToString(bData)

	return &strData, nil
}
