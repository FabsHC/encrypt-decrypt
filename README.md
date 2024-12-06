# go-template-repository
![go-template-repository](https://img.shields.io/badge/go--template--repository-gray?logo=go)
![technology Go 1.23](https://img.shields.io/badge/technology-go%201.23-blue.svg)
![Build & Test](https://github.com/FabsHC/go-template-repository/actions/workflows/go-ci.yml/badge.svg)
[![Go Coverage](https://github.com/FabsHC/go-template-repository/wiki/coverage.svg)](https://raw.githack.com/wiki/FabsHC/go-template-repository/coverage.html)

A template repository for Golang.

To use it, change the module in [go.mod](go.mod) file and update the badges URL to your current project.

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