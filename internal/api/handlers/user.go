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
	auth *services.AuthService
}

func NewUserHandler(svc *services.UserService, auth *services.AuthService) *UserHandler {
    return &UserHandler{svc: svc, auth: auth}
}


func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, errStack := h.svc.GetAllUsers()
	if returnHttpError(w, errStack) { return }
	returnSuccess(w, users)
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

	// Send Verification Email
	httpStack = h.auth.SendEmailVerification(email.String(), username.String())
	if returnHttpError(w, httpStack) { return }

	// Return Success
	returnSuccess(w, "User Registered Successfully")
}

func (h *UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	email, errStack := h.auth.VerifyEmail(token)
	if returnHttpError(w, errStack) { return }

	errStack = h.svc.SetEmailAsVerified(email)
	if returnHttpError(w, errStack) { return }

	returnSuccess(w, "Email Verified")
}