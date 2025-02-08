package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) SerieHandler(w http.ResponseWriter, r *http.Request) {
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
