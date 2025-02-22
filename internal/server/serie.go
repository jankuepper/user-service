package server

import (
	"auth-service/internal/database"
	"encoding/json"
	"net/http"
)

func (s *Server) SerieHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSeries(w, s)
	case http.MethodPost:
		postSeries(w, r, s)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getSeries(w http.ResponseWriter, s *Server) {
	resp := make(map[string]any)
	series, err := s.db.GetAllSeries()
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	resp["success"] = true
	resp["errors"] = []string{}
	resp["series"] = series
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}

func postSeries(w http.ResponseWriter, r *http.Request, s *Server) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var data database.SerieData
	err := decoder.Decode(&data)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	_, err = s.db.CreateSerie(data)
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
