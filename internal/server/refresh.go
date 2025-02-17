package server

import (
	"auth-service/internal/services"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		refreshToken(w, r)
	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	authHeader := r.Header.Get("Authorization")
	token := strings.ReplaceAll(authHeader, "Bearer ", "")
	if token == "" {
		token = r.URL.Query().Get("jwt")
	}
	jwtToken, err := services.VerifyToken(token)
	if err != nil {
		http.Error(w, "Please sign in", http.StatusUnauthorized)
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userEmail := claims["email"].(string)
	token, err = services.CreateToken(userEmail)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	resp["success"] = true
	resp["errors"] = []string{}
	resp["token"] = token
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
