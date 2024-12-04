package handlers

import (
	"net/http"
	"time"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/services"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

type UserHandler struct {
    svc *services.UserService
	auth *services.AuthService
}

func NewUserHandler(svc *services.UserService, auth *services.AuthService) *UserHandler {
    return &UserHandler{svc: svc, auth: auth}
}


func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get All Users
	httpStack := errpkg.NewHttpStack().WithStatus(http.StatusBadRequest)
	users := h.svc.GetAllUsers(httpStack)
	if returnHttpError(w, httpStack) { return }

	// Return Success
	returnSuccess(w, users)
}


func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {	
	// Get Form Data
	httpStack := errpkg.NewHttpStack().WithStatus(http.StatusBadRequest)
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
	httpStack.WithStatus(http.StatusConflict)
	h.svc.RegisterUser(user, httpStack)
	if returnHttpError(w, httpStack) { return }

	// Send Verification Email
	httpStack.WithStatus(http.StatusInternalServerError)
	h.auth.SendEmailVerification(email.String(), username.String(), httpStack)
	if returnHttpError(w, httpStack) { return }

	// Return Success
	returnSuccess(w, "User Registered Successfully")
}


func (h *UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	// Get Token
	token := r.FormValue("token")

	// Verify Token
	httpStack := errpkg.NewHttpStack().WithStatus(http.StatusBadRequest)
	email := h.auth.VerifyEmail(token, httpStack)
	if returnHttpError(w, httpStack) { return }

	// Update User Details
	httpStack.WithStatus(http.StatusInternalServerError)
	h.svc.SetEmailAsVerified(email, httpStack)
	if returnHttpError(w, httpStack) { return }

	// Return Success
	returnSuccess(w, "Email Verified")
}


func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Get Form Data
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Create Login Model
	loginInfo := models.NewLogin(
		email,
		username,
		password,
	)

	// Attempt Login
	httpStack := errpkg.NewHttpStack().WithStatus(http.StatusBadRequest)
	userID, username := h.svc.AttemptLogin(*loginInfo, httpStack)
	if returnHttpError(w, httpStack) { return }

	// Check For Existing Session
	httpStack.WithStatus(http.StatusForbidden)
	structuredErr := errs.LoginAlreadyLoggedIn("You are already logged in.")
	_, err := r.Cookie("session_id")
	if err == nil { httpStack.Add(structuredErr) }
	if returnHttpError(w, httpStack) { return }

	// Create Session
	httpStack.WithStatus(http.StatusInternalServerError)
	session := h.auth.CreateSession(userID, httpStack)
	if returnHttpError(w, httpStack) { return }
	
	// Set Session Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(session.ExpiryInNS)),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	// Set CSRF Token Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    session.Data.CSRFToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(session.ExpiryInNS)),
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	
	// Return Success
	returnSuccess(w, "Welcome back, " + username)
}


func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Get Cookie
	httpStack := errpkg.NewHttpStack().WithStatus(http.StatusInternalServerError)
	sessionCookie, err := r.Cookie("session_id")
	if err != nil { errs.Internal(httpStack, err) }
	if returnHttpError(w, httpStack) { return }

	// Delete Session
	h.auth.DeleteSession(sessionCookie.Value, httpStack)
	if returnHttpError(w, httpStack) { return }

	// Set Session Cookie To Blank
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})

	// Set CSRF Token Cookie To Blank
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})
	
	// Return Success
	returnSuccess(w, "You have been successfully logged out.")
}