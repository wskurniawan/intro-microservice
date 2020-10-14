package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wskurniawan/intro-microservice/service-product/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-product", http.HandlerFunc(handler.AddMenuHandler))

	fmt.Println("Server listen on :8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
