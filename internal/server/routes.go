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
	mux.HandleFunc("/", s.HelloWorldHandler)
	mux.Handle("/health", middleware.Auth(http.HandlerFunc(s.healthHandler)))

	dir := os.Getenv("DATA_DIR")
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", dir)
		return mux
	}
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/data/", http.StripPrefix("/data/", middleware.Auth(fs)))

	mux.HandleFunc("/signup", s.SignUpHandler)
	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
