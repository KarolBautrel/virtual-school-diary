package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"virtual-diary/middlewares"

	"github.com/gorilla/mux"
)

type StudentInput struct {
	Name    string `json:"name"`
	Age     string `json:"profile"`
	Surname string `json:"surname"`
	ClassId string `json:"classid"`
}

func RegisterRoutes(router *mux.Router, readService *StudentReadService, writeService *StudentWriteService) {
	router.HandleFunc("/student", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			StudentHandler(w, r, readService)
		case "POST":
			StudentCreateHandler(w, r, writeService)
		case "DELETE":
			StudentDeleteHandler(w, r, writeService)

		}
	}))

	router.HandleFunc("/student/{id}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		StudentByIdHandler(w, r, readService)
	})).Methods("GET")

	router.HandleFunc("/student/class/{class_id}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		StudentsByClassHandler(w, r, readService)
	}))
}

func StudentHandler(w http.ResponseWriter, r *http.Request, readService *StudentReadService) {
	message := readService.GetWelcomeMessage()
	fmt.Fprintln(w, message)
}

func StudentByIdHandler(w http.ResponseWriter, r *http.Request, readService *StudentReadService) {
	vars := mux.Vars(r)
	id := vars["id"]
	mes, err := readService.GetStudentById(id)
	if err != nil {
		http.Error(w, "student Not Found", http.StatusNotFound)
		return

	}
	fmt.Fprintln(w, mes)
}

func StudentsByClassHandler(w http.ResponseWriter, r *http.Request, readService *StudentReadService) {
	vars := mux.Vars(r)
	class_id := vars["class_id"]
	mes, err := readService.GetStudentsByClass(class_id)
	if err != nil {
		http.Error(w, "student Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, mes)
}

func StudentDeleteHandler(w http.ResponseWriter, r *http.Request, writeService *StudentWriteService) {
	vars := mux.Vars(r)
	id := vars["id"]
	mes, err := writeService.DeleteStudent(id)
	if err != nil {
		http.Error(w, "There was an error with student deletion", http.StatusBadGateway)
		return

	}
	fmt.Fprintln(w, mes)
}

func StudentCreateHandler(w http.ResponseWriter, r *http.Request, writeService *StudentWriteService) {
	studnetSchema := StudentInput{}
	if err := json.NewDecoder(r.Body).Decode(&studnetSchema); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	status, err := writeService.CreateStudent(studnetSchema.Name, studnetSchema.Surname, studnetSchema.Age, studnetSchema.ClassId)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadGateway)
		return
	}
	fmt.Fprintln(w, status)

}
