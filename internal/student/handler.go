package student

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes rejestruje trasy dla operacji związanych z klasami.
func RegisterRoutes(router *mux.Router, service StudentService) {
	router.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		ClassHandler(w, r, service)
	}).Methods("GET")
}

// ClassHandler obsługuje zapytania GET dla strony klasy.
func ClassHandler(w http.ResponseWriter, r *http.Request, service StudentService) {
	message := service.GetWelcomeMessage()
	fmt.Fprintln(w, message)
}
