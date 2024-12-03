package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/kye-gregory/koicards-api/internal/models"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type SessionStore struct {
	db *redis.Client
}

func NewSessionStore(db *redis.Client) *SessionStore {
	return &SessionStore{db: db}
}

func (store *SessionStore) CreateSession(userID userVO.ID) (*models.Session, error) {return nil, fmt.Errorf("Not Implemented")}
func (store *SessionStore) DeleteSession(sessionID string) error {return fmt.Errorf("Not Implemented")}
func (store *SessionStore) VerifySession(sessionID string) (bool, error) {return false, fmt.Errorf("Not Implemented")}