package main

import (
	"log"
	"net/http"
	"randomuser/handlers"
)

func main() {
	http.HandleFunc("/randomuser", handlers.RandomUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
