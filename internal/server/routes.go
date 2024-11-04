package server

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /register", 	handleUserRegister)
}

func handleUserRegister (w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate Form Data
	if (len(username)<8 || len(password)<8) {
		err := http.StatusAccepted
		http.Error(w, "invalid username or password", err)
		return
	}

	// Check If User Exists
	if _, ok := users[username]; ok {
		err := http.StatusConflict
		http.Error(w, "user already exists", err)
		return
	}

	hashedPassword, _ := hash(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Registered Successfully!")
}