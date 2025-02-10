package server

import (
	"auth-service/internal/database"
	"encoding/json"
	"net/http"
)

func (s *Server) SeasonHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSeason(w, s)
	case http.MethodPost:
		postSeason(w, r, s)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getSeason(w http.ResponseWriter, s *Server) {
	resp := make(map[string]any)
	seasons, err := s.db.GetAllSeasons()
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	resp["success"] = true
	resp["errors"] = []string{}
	resp["seasons"] = seasons
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}

func postSeason(w http.ResponseWriter, r *http.Request, s *Server) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var data database.SeasonData
	err := decoder.Decode(&data)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	_, err = s.db.CreateSeason(data)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	resp["success"] = true
	resp["errors"] = []string{}
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
