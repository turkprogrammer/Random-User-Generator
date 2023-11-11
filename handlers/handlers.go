package handlers

import (
	"html/template"
	"log"
	"net/http"
	"randomuser/user"
)

func RandomUserHandler(w http.ResponseWriter, r *http.Request) {
	repo := user.NewRandomUserRepository("https://randomuser.me/api/")
	interactor := user.NewRandomUserInteractor(repo)

	randomUser, err := interactor.GetRandomUser()
	if err != nil {
		log.Printf("Error getting random user: %v\n", err)
		http.Error(w, "Error getting random user", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("user.html").ParseFiles("templates/user.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, randomUser)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
