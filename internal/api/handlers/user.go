package handlers

import (
	"fmt"
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/services"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {	
	// Create User Struct
	user := models.User {
		Email: 		r.FormValue("email"),
		Username: 	r.FormValue("username"),
		Password: 	r.FormValue("password"),
	}

	// Validate
	err := h.service.ValidateUser(&user)
	if ReturnHttpError(w, err, http.StatusBadRequest) { return }

	// Register
	err = h.service.RegisterUser(&user)
	if ReturnHttpError(w, err, http.StatusConflict) { return }

	// Return Success
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Registered Successfully!")
}
