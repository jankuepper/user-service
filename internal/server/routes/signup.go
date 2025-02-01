package server

import (
	"auth-service/internal/database"
	"encoding/json"
	"log"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var test database.UserData
	err := decoder.Decode(&test)
	if err != nil {
		log.Printf("An error occured during sign-up.")
		resp["success"] = false
	} else {

	}
	jsonResp, err := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
