# encrypt-decrypt
![encrypt-decrypt](https://img.shields.io/badge/encrypt--decrypt-gray?logo=go)
![technology Go 1.23](https://img.shields.io/badge/technology-go%201.23-blue.svg)
![Build & Test](https://github.com/FabsHC/encrypt-decrypt/actions/workflows/go-ci.yml/badge.svg)
[![Go Coverage](https://github.com/FabsHC/encrypt-decrypt/wiki/coverage.svg)](https://raw.githack.com/wiki/FabsHC/encrypt-decrypt/coverage.html)

A personal study project in Go (v1.23) that demonstrates how to generate and manage HMAC keys, encrypt and decrypt messages using AES-GCM, and store data securely in a lightweight embedded database (`buntdb`).

## Features

- **Key Management**:
    - Generate and store an HMAC key.
    - Retrieve the stored key via an API endpoint.
- **Encryption/Decryption**:
    - Encrypt a message using AES-GCM.
    - Decrypt a previously encrypted message.
- **Lightweight Database**:
    - Stores the HMAC key using [`buntdb`](https://github.com/tidwall/buntdb).

## Endpoints

### 1. `POST /key`
- **Description**: Generates a new HMAC key and stores it in the database.
- **Response**: 201 Created
  ```json
  {
  "key": "0766ff57136f0d93328d990d57404b7dfbdac4a5fe350bbf8d3e9f108366599e"
  }
  ```

### 2. `GET /key`
- **Description**: Retrieves the stored HMAC key.
- **Response**:
  ```json
  {
    "key": "a1b2c3d4e5f67890abcdef1234567890abcdef1234567890abcdef1234567890"
  }
  ```

### 3. `POST /encrypt`
- **Description**: Encrypts a given plaintext using AES-GCM.
- **Request**:
  ```json
  {
    "text": "fabs"
  }
  ```
- **Response**:
  ```json
  {
    "text_decrypted": "fabs",
    "text_encrypted": "kIQUVGVxJOGsSB2imBc48PJ+RcUg/UEteMX+0qberDM="
  }
  ```

### 4. `POST /decrypt`
- **Description**: Decrypts a previously encrypted message using AES-GCM.
- **Request**:
  ```json
  {
    "text": "kIQUVGVxJOGsSB2imBc48PJ+RcUg/UEteMX+0qberDM="
  }
  ```
- **Response**:
  ```json
  {
    "text_decrypted": "fabs",
    "text_encrypted": "kIQUVGVxJOGsSB2imBc48PJ+RcUg/UEteMX+0qberDM="
  }
  ```

## Project Structure

```
encrypt-decrypt/
├── cmd/
│   └── main.go          # Arquivo principal para iniciar o servidor
├── internal/
│   ├── handler/         # Camada responsável por HTTP (controllers)
│   │   └── user_handler.go
│   ├── service/         # Camada de regras de negócio (use cases)
│   │   └── user_service.go
│   ├── repository/      # Camada de acesso a dados
│   │   └── user_repo.go
│   └── entity/          # Camada de entidades de domínio
│       └── user.go
├── go.mod               # Gerenciamento de dependências
└── go.sum
```

## Requirements

- Go 1.23 or later
- buntdb library



## Installation

### 1. Clone repository
```bash
git clone https://github.com/your-username/encrypt-decrypt.git
cd encrypt-decrypt
```

### 2. Install dependencies
```bash
go mod tidy
```

### 3. Run the application
```bash
go run cmd/main.go
```

## How it works

### 1. Key Management:
- The POST /key endpoint generates a cryptographically secure HMAC key and stores it in buntdb.
- The GET /key endpoint retrieves this key for internal operations.

### 2. Encryption/Decryption:
- AES-GCM is used for encrypting and decrypting messages. This ensures authenticated encryption for message confidentiality and integrity.
- Encrypted messages are encoded in Base64 for easy handling.

### 3. Database:
- buntdb is used to persist the HMAC key. It provides a lightweight, in-memory database with ACID compliance



## Example Usage
### 1. Generate a key:
```bash
curl -X POST http://localhost:8080/key
```

### 2.Retrieve the key:
```bash
curl http://localhost:8080/key
```

### 3. Encrypt a message:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"text":"fabs"}' http://localhost:8080/encrypt
```

### 4.Decrypt a message:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"text":"fabs"}' http://localhost:8080/decrypt
```

## Author
Developed by Fabs.