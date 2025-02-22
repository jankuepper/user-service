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

	// public
	mux.Handle("/", middleware.Cors(http.HandlerFunc(s.LoginHandler)))
	mux.Handle("/signup", middleware.Cors(http.HandlerFunc(s.SignUpHandler)))
	mux.Handle("/health", middleware.Cors(http.HandlerFunc(s.healthHandler)))

	// private
	fs := getFileServer()
	// mux.Handle("/upload", middleware.Cors(middleware.Auth(http.HandlerFunc(s.FileUploadHandler)))) todo: fix
	mux.Handle("/refresh", middleware.Cors(middleware.Auth(http.HandlerFunc(s.RefreshHandler))))
	mux.Handle("/data/", http.StripPrefix("/data/", middleware.Cors(middleware.Auth(fs))))
	mux.Handle("/series", middleware.Cors(middleware.Auth(http.HandlerFunc(s.SerieHandler))))
	mux.Handle("/seasons", middleware.Cors(middleware.Auth(http.HandlerFunc(s.SeasonHandler))))
	mux.Handle("/episodes", middleware.Cors(middleware.Auth(http.HandlerFunc(s.EpisodeHandler))))
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
