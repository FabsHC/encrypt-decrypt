package main

import (
	"encrypt-decrypt/internal/handler"
	"encrypt-decrypt/internal/repository"
	"encrypt-decrypt/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	dbPath := "key.db"
	keyRepo, err := repository.NewKeyRepository(dbPath)
	if err != nil {
		log.Fatalf("fail to initialize database: %v", err)
	}
	defer keyRepo.Close()

	keyService := service.NewKeyService(keyRepo)
	keyHandler := handler.NewKeyHandler(keyService)

	encryptService := service.NewEncryptService(keyRepo)
	encryptHandler := handler.NewEncryptHandler(encryptService)

	decryptService := service.NewDecryptService(keyRepo)
	decryptHandler := handler.NewDecryptHandler(decryptService)

	r := mux.NewRouter()
	r.HandleFunc("/key", keyHandler.CreateNewKey).Methods("POST")
	r.HandleFunc("/key", keyHandler.GetKey).Methods("GET")

	r.HandleFunc("/encrypt", encryptHandler.EncryptData).Methods("POST")
	r.HandleFunc("/decrypt", decryptHandler.DecryptData).Methods("POST")

	log.Println("server listen to port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
