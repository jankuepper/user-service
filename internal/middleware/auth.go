package middleware

import (
	"auth-service/internal/services"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.ReplaceAll(authHeader, "Bearer ", "")
		if token == "" {
			token = r.URL.Query().Get("jwt")
		}
		token, err := services.VerifyToken(token)
		if err != nil {
			http.Error(w, "Please sign in", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
