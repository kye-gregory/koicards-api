package server

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{userID}", GetUserIDHandler)
}

func GetUserIDHandler (w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("User ID:" + userID))
}