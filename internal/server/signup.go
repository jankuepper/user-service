package server

import (
	"auth-service/internal/database"
	"auth-service/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var userData database.UserData
	err := decoder.Decode(&userData)
	if err != nil {
		log.Printf("An error occured during sign-up.")
		resp["success"] = false
		resp["errors"] = []string{"Something went wrong."}
		resp["token"] = nil
		return
	} else {
		s.db.CreateUser(userData)
		token, _ := services.CreateToken(userData.Email)
		resp["success"] = true
		resp["errors"] = []string{}
		resp["token"] = token
	}
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
