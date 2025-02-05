package server

import (
	"auth-service/internal/database"
	"auth-service/internal/services"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type loginUserData struct {
	Email    string
	Password string
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var userData loginUserData
	err := decoder.Decode(&userData)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	var res database.User
	res, err = s.db.GetUserByEmail(userData.Email)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Data.Password), []byte(userData.Password))
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
