package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/kye-gregory/koicards-api/internal/models"
)

type SessionStore struct {
	db *redis.Client
}

func NewSessionStore(db *redis.Client) *SessionStore {
	return &SessionStore{db: db}
}

func (store *SessionStore) CreateSession(userID int) (*models.Session, error) {return nil, fmt.Errorf("Not Implemented")}
func (store *SessionStore) DeleteSession(sessionID string) error {return fmt.Errorf("Not Implemented")}
func (store *SessionStore) VerifySession(sessionID string) (bool, error) {return false, fmt.Errorf("Not Implemented")}