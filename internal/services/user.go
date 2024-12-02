package services

import (
	"github.com/kye-gregory/koicards-api/pkg/auth"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

type UserService struct {
	store 	store.UserStore
}

// Constructor function for UserService
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}


func (svc *UserService) GetAllUsers(errStack *errpkg.HttpStack) ([]*models.User) {
	// Get All Users
	users, err := svc.store.GetAllUsers()
	if err != nil { errs.Internal(errStack, err); return nil }

	// Return
	return users
}


// Calls store to register database
func (svc *UserService) RegisterUser(u *models.User, errStack *errpkg.HttpStack) {
	// Check If Email Is Already Registered
	exists, err := svc.store.IsEmailRegistered(u.Email.String())
	if err != nil { errs.Internal(errStack, err); return }
	if exists { errStack.Add(errs.EmailInUse("email already in use")) }

	// Check If Username Is Already Registered
	exists, err = svc.store.IsUsernameRegistered(u.Username.String())
	if err != nil { errs.Internal(errStack, err); return }
	if exists { errStack.Add(errs.UsernameInUse("username already in use")) }

	// Return Non-Internal Errors
	if !errStack.IsEmpty() { return }

	// Update Store
	err = svc.store.CreateUser(u)
	if err != nil { errs.Internal(errStack, err); return }
}


func (svc *UserService) SetEmailAsVerified(email string, errStack *errpkg.HttpStack) {
	// Update Store
	err := svc.store.ActivateUser(email)
	if err != nil { errs.Internal(errStack, err); return }
}


func (svc *UserService) AttemptLogin(loginInfo models.Login, errStack *errpkg.HttpStack) *userVO.ID {
	// Initialise Error
	structuredErr := errs.LoginInvalidDetails("incorrect username or password")

	// Get User
	user, err := svc.store.GetUser(loginInfo.Identifier)
	if err != nil { errs.Internal(errStack, err); return nil }
	if user == nil { errStack.Add(structuredErr); return nil }

	// Check Password
	isUser := auth.VerifyPassword(user.Password.String(), loginInfo.Password)
	if ( !isUser ) { errStack.Add(structuredErr); return nil }

	return &user.ID
}