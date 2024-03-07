package class

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ClassInput struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

func RegisterRoutes(router *mux.Router, readService ClassReadService) {
	router.HandleFunc("/class", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			AllClasses(w, r, readService)
		}
	}).Methods("GET", "POST", "DELETE")

	router.HandleFunc("/class/{id}", func(w http.ResponseWriter, r *http.Request) {
		ClassByIdHandler(w, r, readService)
	}).Methods("GET")
}

func AllClasses(w http.ResponseWriter, r *http.Request, readService ClassReadService) {
	classes, err := readService.GetAllClasses()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(classes)

	fmt.Fprintln(w, classes)
}
func ClassByIdHandler(w http.ResponseWriter, r *http.Request, readService ClassReadService) {
	vars := mux.Vars(r)
	id := vars["id"]
	class, err := readService.GetClassById(id)
	if err != nil {
		http.Error(w, "Class not found", http.StatusNotFound)

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(class)

	fmt.Fprintln(w, class)
}
