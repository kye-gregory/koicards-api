package services

import (
	"fmt"

	"github.com/kye-gregory/koicards-api/internal/auth"
	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
)

type UserService struct {
	store 	store.UserStore
}

// Constructor function for AuthService
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}


// Checks user details are in a valid format
func (s *UserService) ValidateUser(user *models.User) error {
	return nil
}


// Calls store to register database
func (s *UserService) RegisterUser(user *models.User) error {
	// Check if user exists
	exists, err := s.store.UserExists(user.Email)
	if err != nil { return err }
	if exists { return fmt.Errorf("user already exists") }

	// Has password
	hashedPassword, err := auth.Hash(user.Password)
	if err != nil { return err }
	user.Password = hashedPassword

	// Update store & return any errors
	return s.store.CreateUser(user)
}