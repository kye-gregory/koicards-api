package services

import (
	"errors"

	"github.com/kye-gregory/koicards-api/internal/auth"
	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
	e "github.com/kye-gregory/koicards-api/pkg/errors"
)

type UserService struct {
	store 	store.UserStore
}

// Constructor function for UserService
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}

func (s *UserService) ValidateUser(user *models.User, status int) *e.HttpErrorStack {
	// Create Error Stack
	errStack := e.NewHttpErrorStack(status)

	// Validate Through Auth Package
	auth.ValidateEmail(errStack, user.Email)
	auth.ValidateUsername(errStack, user.Username)
	auth.ValidatePassword(errStack, user.Password)

	// Return Error Stack
	return errStack
}


// Calls store to register database
func (s *UserService) RegisterUser(u *models.User, status int) *e.HttpErrorStack {
	// Create Error Stack
	errStack := e.NewHttpErrorStack(status)

	// Check If Email Is Already Registered
	exists, err := s.store.IsEmailRegistered(u.Email)
	if err != nil { return errStack.ReturnInternalError() }
	if exists { 
		err = errors.New("email already in use")
		errStack.Add("database", err.Error())
	}

	// Check If Username Is Already Registered
	exists, err = s.store.IsUsernameRegistered(u.Username)
	if err != nil { return errStack.ReturnInternalError() }
	if exists { 
		err = errors.New("username already taken")
		errStack.Add("database", err.Error())
	}

	// Return Non-Internal Errors
	if !errStack.IsEmpty() { return errStack }

	// Hash Password
	hashedPassword, err := auth.Hash(u.Password)
	if err != nil { return errStack.ReturnInternalError() }
	u.Password = hashedPassword

	// Update Store
	err = s.store.CreateUser(u)
	if (err != nil ) { return errStack.ReturnInternalError() }

	// Return Error Stack
	return errStack
}