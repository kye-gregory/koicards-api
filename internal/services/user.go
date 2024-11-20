package services

import (
	"errors"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

type UserService struct {
	store 	store.UserStore
}

// Constructor function for UserService
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}



// Calls store to register database
func (svc *UserService) RegisterUser(u *models.User, status int) *errorstack.HttpStack {
	// Create Error Stack
	errStack := errorstack.NewHttpStack().Status(status)

	// Check If Email Is Already Registered
	exists, err := svc.store.IsEmailRegistered(u.Email.String())
	if err != nil { return errStack.ReturnInternalError() }
	if exists { 
		err = errors.New("email already in use")
		errStack.Add("database", err)
	}

	// Check If Username Is Already Registered
	exists, err = svc.store.IsUsernameRegistered(u.Username.String())
	if err != nil { return errStack.ReturnInternalError() }
	if exists { 
		err = errors.New("username already taken")
		errStack.Add("database", err)
	}

	// Return Non-Internal Errors
	if !errStack.IsEmpty() { return errStack }

	// Update Store
	err = svc.store.CreateUser(u)
	if (err != nil ) { return errStack.ReturnInternalError() }

	// Return Error Stack
	return errStack
}


func (svc *UserService) SetEmailAsVerified(email string) *errorstack.HttpStack {
	// Create Error Stack
	errStack := errorstack.NewHttpStack()

	err := svc.store.ActivateUser(email)
	if err != nil { return errStack.ReturnInternalError() }

	return errStack
}