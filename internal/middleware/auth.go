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
		err := services.VerifyToken(token)
		if err != nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
