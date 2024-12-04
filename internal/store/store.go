package store

import (
	"github.com/kye-gregory/koicards-api/internal/models"
)

type Database struct {
    UserStore    UserStore
    SessionStore SessionStore
}

// NewDatabase initializes the database with its stores
func NewDatabase(userStore UserStore, sessionStore SessionStore) *Database {
    return &Database{
        UserStore:    userStore,
        SessionStore: sessionStore,
    }
}

// UserStore defines the methods for user data operations
type UserStore interface {
	IsUsernameRegistered(email string) (bool, error)
    IsEmailRegistered(email string) (bool, error)
    CreateUser(user *models.User) error
    VerifyEmail(email string) error 
    GetUserByEmail(identifier string) (*models.User, error)
    GetUserByUsername(identifier string) (*models.User, error)
    GetAllUsers() ([]*models.User, error)
}


type SessionStore interface {
    CreateSession(userID int) (*models.Session, error)
    DeleteSession(sessionID string) error
    VerifySession(sessionID string) (bool, error)
}