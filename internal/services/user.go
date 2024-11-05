package services

import (
	"fmt"

	"github.com/kye-gregory/koicards-api/internal/models"
	"github.com/kye-gregory/koicards-api/internal/store"
)

type UserService struct {
	store store.UserStore
}

// Constructor function for AuthService
func NewUserService(s store.UserStore) *UserService {
	return &UserService{store: s}
}

// Check if a user already exists in the database
func (s *UserService) RegisterUser(user *models.User) error {
	exists, err := s.store.UserExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	return s.store.CreateUser(user)
}