package server

import (
	"io/ioutil"
	"net/http"
	"os"
)

func (s *Server) FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodPost:
		fileUpload(r)
	case http.MethodPut:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	case http.MethodDelete:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func fileUpload(r *http.Request) {
	dir := os.Getenv("DATA_DIR")
	r.ParseMultipartForm(1000)
	file, _, err := r.FormFile("test.mp4")
	if err != nil {
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(dir, "upload-*.mp4")
	if err != nil {
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	tempFile.Write(fileBytes)
}
