package server

import (
	"log"
	"net/http"
)


func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("Method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w,r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "ValidAuthToken" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w,r)
	}
}