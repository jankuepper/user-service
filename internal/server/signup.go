package server

import (
	"auth-service/internal/database"
	"auth-service/internal/services"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var userData database.UserData
	err := decoder.Decode(&userData)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	userData.Password = string(hashedPassword)
	_, err = s.db.CreateUser(userData)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	token := ""
	token, err = services.CreateToken(userData.Email)
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

func returnError(err error, resp map[string]any) []byte {
	log.Print("An error occured during sign-up. ", err)
	resp["success"] = false
	resp["errors"] = []string{"Something went wrong."}
	resp["token"] = nil
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Print("An error occured during sign-up error response. ", err)
	}
	return jsonResp
}
