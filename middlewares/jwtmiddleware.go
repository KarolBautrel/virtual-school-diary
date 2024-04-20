package middlewares

import (
	"net/http"
)

func JWTmiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := VerifyToken(token); err != nil {
			http.Error(w, "jwt verify error", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
