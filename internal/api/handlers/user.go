package handlers

import (
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/services"
	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
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
	errStack := h.service.ValidateUser(&user, http.StatusBadRequest)
	if returnHttpError(w, errStack) { return }

	// Register
	errStack = h.service.RegisterUser(&user, http.StatusConflict)
	if returnHttpError(w, errStack) { return }

	// Return Success
	returnTextSuccess(w, "User Registered Successfully!")
}
