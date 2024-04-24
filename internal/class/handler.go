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
	})).Methods("GET", "POST", "DELETE")

	router.HandleFunc("/class/{id}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		ClassByIdHandler(w, r, readService)
	})).Methods("GET")

	router.HandleFunc("/class/{classId}/remove-student/{studentId}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		DeleteStudentFromClassHandler(w, r, writeService)
	})).Methods("DELETE")

	router.HandleFunc("/class/{classId}/add-student/{studentId}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		AddStudentToClassHandler(w, r, writeService)
	})).Methods("DELETE")
}

func AllClasses(w http.ResponseWriter, r *http.Request, readService *ClassReadService) {
	classes, err := readService.GetAllClasses()
	if err != nil {
		http.Error(w, "error with fetching classes", http.StatusBadRequest)
		return

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
		http.Error(w, "error with fetching class", http.StatusNotFound)
		return
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
		http.Error(w, "error with create class", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)
}

func RemoveClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	vars := mux.Vars(r)
	classId, studentId := vars["classId"], vars["studentId"]
	status, error := writeService.RemoveStudentFromClass(classId, studentId)
	if error != nil {
		http.Error(w, "error with removing class", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)
}

func DeleteStudentFromClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	status, error := writeService.RemoveClass(studentId)
	if error != nil {
		http.Error(w, "error with removing student from class", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)
}
func AddStudentToClassHandler(w http.ResponseWriter, r *http.Request, writeService *ClassWriteService) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	classId := vars["classId"]
	if classId == "" || studentId == "" {
		http.Error(w, "error with removing student from class", http.StatusBadGateway)
		return
	}
	status, error := writeService.AddStudentToClass(studentId, classId)
	if error != nil {
		http.Error(w, "error with removing student from class", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)
}
