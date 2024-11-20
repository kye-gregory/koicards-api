package store

import "github.com/kye-gregory/koicards-api/internal/models"

type Database struct {
    UserStore    UserStore
}

// NewDatabase initializes the database with its stores
func NewDatabase(userStore UserStore) *Database {
    return &Database{
        UserStore:    userStore,
    }
}

// UserStore defines the methods for user data operations
type UserStore interface {
	IsUsernameRegistered(email string) (bool, error)
    IsEmailRegistered(email string) (bool, error)
    CreateUser(user *models.User) error
    ActivateUser(email string) error 
    GetUserByEmail(email string) (*models.User, error)
    GetAllUsers() ([]*models.User, error)
}