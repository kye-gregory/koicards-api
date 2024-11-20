package handlers

import (
	"net/http"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/services"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

type UserHandler struct {
    svc *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {	
	// Create User Struct
	httpStack := errorstack.NewHttpStack().Status(http.StatusBadRequest)
	email := userVO.NewEmail(r.FormValue("email"), httpStack)
	username := userVO.NewUsername(r.FormValue("username"), httpStack)
	password := userVO.NewPassword(r.FormValue("password"), httpStack)
	if returnHttpError(w, httpStack) { return }

	// Create User Model
	user := models.NewUser (
		*email,
		*username,
		*password,
	)

	// Register
	httpStack = h.svc.RegisterUser(user, http.StatusConflict)
	if returnHttpError(w, httpStack) { return }

	// Return Success
	returnTextSuccess(w, "User Registered Successfully!")
}
