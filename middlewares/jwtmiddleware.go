package middlewares

import (
	"fmt"
	"net/http"
)

func JWTmiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := VerifyToken(token); err != nil {
			fmt.Errorf("jwt verify error")
		}
		next.ServeHTTP(w, r)
	}
}
