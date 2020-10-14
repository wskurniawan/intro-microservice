package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wskurniawan/intro-microservice/auth/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
