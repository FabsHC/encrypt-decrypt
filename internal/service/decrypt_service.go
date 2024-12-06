package service

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"encrypt-decrypt/internal/repository"
)

type (
	DecryptService interface {
		Decrypt(encryptedText string) (*string, error)
	}

	decryptService struct {
		repo repository.KeyRepository
	}
)

func NewDecryptService(repo repository.KeyRepository) DecryptService {
	return &decryptService{
		repo: repo,
	}
}

func (d *decryptService) Decrypt(encryptedText string) (*string, error) {
	// Decodifica o texto criptografado de Base64 para bytes
	bEncryptedText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}

	// Obtém a chave do repositório
	key, err := d.repo.Get()
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

	// Divide o nonce e os dados criptografados
	nonceSize := gcm.NonceSize()
	if len(bEncryptedText) < nonceSize {
		return nil, err
	}
	nonce, bEncryptedText := bEncryptedText[:nonceSize], bEncryptedText[nonceSize:]

	// Descriptografa os dados
	bData, err := gcm.Open(nil, nonce, bEncryptedText, nil)
	if err != nil {
		return nil, err
	}

	// Converte os dados descriptografados para string
	data := string(bData)

	return &data, nil
}
