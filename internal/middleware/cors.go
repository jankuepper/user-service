package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// look at the requerst origin from r and then make a decision which origin to  set in the cors header
		// https://stackoverflow.com/questions/47298310/golang-http-allow-certain-domain-name-both-with-www-and-without
		switch origin := r.Header.Get("Origin"); origin {
		case "www.jankuepper.de", "jankuepper.de":
			w.Header().Add("Access-Control-Allow-Origin", origin)
		case "www.jankuepper.eu", "jankuepper.eu":
			w.Header().Add("Access-Control-Allow-Origin", origin)
		case "localhost:5173", "localhost:4173":
			w.Header().Add("Access-Control-Allow-Origin", origin)
		default:
			w.Header().Add("Access-Control-Allow-Origin", "www.jankuepper.eu")
		}
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
