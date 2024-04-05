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
	var siginInInput SignInInput
	if err := json.NewDecoder(r.Body).Decode(&siginInInput); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := readService.SignIn(siginInInput.Username, siginInInput.Password)

	if err != nil {
		http.Error(w, "something went wrong", http.StatusBadRequest)
		return

	}

	fmt.Fprintln(w, token)

}

func GetUserByUsername(w http.ResponseWriter, r *http.Request, readService *AuthReadService) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := readService.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "something went wrong wit getting user by username", http.StatusBadRequest)
		return

	}
	fmt.Fprintln(w, user)
}
func GetUserById(w http.ResponseWriter, r *http.Request, readService *AuthReadService) {
	vars := mux.Vars(r)
	userId := vars["id"]
	user, err := readService.GetUserById(userId)
	if err != nil {
		http.Error(w, "something went wrong wit getting user by id", http.StatusBadRequest)
		return

	}
	fmt.Fprintln(w, user)
}
