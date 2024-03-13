package class

import (
	"encoding/json"
	"fmt"
	"net/http"
	"virtual-diary/middlewares"

	"github.com/gorilla/mux"
)

type ClassInput struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

func RegisterRoutes(router *mux.Router, readService *ClassReadService, writeService *ClassWriteService) {
	router.HandleFunc("/class", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			AllClasses(w, r, readService)
		case "POST":
			CreateClassHandler(w, r, writeService)
		case "DELETE":
			RemoveClassHandler(w, r, writeService)
		}
	}))

	router.HandleFunc("/class/{id}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		ClassByIdHandler(w, r, readService)
	})).Methods("GET")

	router.HandleFunc("/class/{classId}/remove-student/{studentId}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		DeleteStudentFromClassHandler(w, r, writeService)
	})).Methods("DELETE")
}

func AllClasses(w http.ResponseWriter, r *http.Request, readService *ClassReadService) {
	classes, err := readService.GetAllClasses()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(classes)

	fmt.Fprintln(w, classes)
}
func ClassByIdHandler(w http.ResponseWriter, r *http.Request, readService *ClassReadService) {
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

func CreateClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	var input ClassInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	status, err := writeService.CreateClass(input.Name, input.Profile)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)
}

func RemoveClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	vars := mux.Vars(r)
	classId, studentId := vars["classId"], vars["studentId"]
	status, error := writeService.RemoveStudentFromClass(classId, studentId)
	if error != nil {
		http.Error(w, "Something went wrong", http.StatusBadGateway)
	}
	fmt.Fprintln(w, status)
}

func DeleteStudentFromClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	status, error := writeService.RemoveClass(studentId)
	if error != nil {
		http.Error(w, "Something went wrong", http.StatusBadGateway)
	}
	fmt.Fprintln(w, status)
}
