package handlers

import (
	"fmt"
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/services"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate Form Data
	if (len(username)<8 || len(password)<8) {
		err := http.StatusBadRequest
		http.Error(w, "invalid username or password", err)
		return
	}

	// Check If User Exists
	// if _, ok := store.Users[username]; ok {
	// 	err := http.StatusConflict
	// 	http.Error(w, "user already exists", err)
	// 	return
	// }

	// Update Storage
	// hashedPassword, _ := auth.Hash(password)
	// storage.Users[username] = models.Login{
	// 	HashedPassword: hashedPassword,
	// }

	// Return Success
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Registered Successfully!")
}
