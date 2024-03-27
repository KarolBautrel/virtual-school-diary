package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"virtual-diary/middlewares"

	"github.com/gorilla/mux"
)

type SignInInput struct {
	Username string `json:"name"`
	Password string `json:"password"`
}
type RegiserInput struct{}

func RegisterRoutes(router *mux.Router, readService *AuthReadService) {
	router.HandleFunc("/auth/signIn", func(w http.ResponseWriter, r *http.Request) {
		SignIn(w, r, readService)
	}).Methods("POST")

	router.HandleFunc("/auth/{username}", middlewares.JWTmiddleware(func(w http.ResponseWriter, r *http.Request) {
		GetUserByUsername(w, r, readService)
	}))
}

func SignIn(w http.ResponseWriter, r *http.Request, readService *AuthReadService) {
	var input SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := readService.SignIn(input.Username, input.Password)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return

	}

	fmt.Fprintln(w, token)

}

func GetUserByUsername(w http.ResponseWriter, r *http.Request, readService *AuthReadService) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := readService.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return

	}
	fmt.Fprintln(w, user)
}
