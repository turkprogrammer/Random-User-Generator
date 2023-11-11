package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"randomuser/handlers"
)

func main() {
	router := mux.NewRouter()
	handlers.SetupRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
