package server

import (
	"auth-service/internal/database"
	"encoding/json"
	"net/http"
)

func (s *Server) EpisodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getEpisode(w, s)
	case http.MethodPost:
		postEpisode(w, r, s)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getEpisode(w http.ResponseWriter, s *Server) {
	resp := make(map[string]any)
	episodes, err := s.db.GetAllEpisodes()
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	resp["success"] = true
	resp["errors"] = []string{}
	resp["episodes"] = episodes
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}

func postEpisode(w http.ResponseWriter, r *http.Request, s *Server) {
	resp := make(map[string]any)
	decoder := json.NewDecoder(r.Body)
	var data database.EpisodeData
	err := decoder.Decode(&data)
	if err != nil {
		res := returnError(err, resp)
		w.Write(res)
		return
	}
	_, err = s.db.CreateEpisode(data)
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
