package class

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, service ClassService) {
	router.HandleFunc("/class", func(w http.ResponseWriter, r *http.Request) {
		ClassHandler(w, r, service)
	}).Methods("GET")

	router.HandleFunc("/class/{id}", func(w http.ResponseWriter, r *http.Request) {
		ClassByIdHandler(w, r, service)
	}).Methods("GET")
}

func ClassHandler(w http.ResponseWriter, r *http.Request, service ClassService) {
	message := service.GetWelcomeMessage()
	fmt.Fprintln(w, message)
}

func ClassByIdHandler(w http.ResponseWriter, r *http.Request, service ClassService) {
	vars := mux.Vars(r)
	id := vars["id"]
	message, err := service.GetClassById(id)
	if err != nil {
		http.Error(w, "Class not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, message)
}
