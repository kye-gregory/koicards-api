package handlers

import (
	"fmt"
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/services"
	e "github.com/kye-gregory/koicards-api/pkg/errors"
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
	if e.ReturnHttpError(w, errStack) { return }

	// Register
	errStack = h.service.RegisterUser(&user, http.StatusConflict)
	if e.ReturnHttpError(w, errStack) { return }

	// Return Success
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Registered Successfully!")
}
