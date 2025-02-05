package server

import (
	"auth-service/internal/middleware"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.LoginHandler)
	mux.HandleFunc("/signup", s.SignUpHandler)
	mux.Handle("/health", middleware.Auth(http.HandlerFunc(s.healthHandler)))
	fs := getFileServer()
	mux.Handle("/data/", http.StripPrefix("/data/", middleware.Auth(fs)))
	return mux
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}

func getFileServer() http.Handler {
	dir := os.Getenv("DATA_DIR")
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(fmt.Sprintf("Directory '%s' not found.\n", dir))
	}
	return http.FileServer(http.Dir(dir))
}
