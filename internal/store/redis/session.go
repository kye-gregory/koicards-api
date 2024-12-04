package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/kye-gregory/koicards-api/internal/models"
)

type SessionStore struct {
	db *redis.Client
}

func NewSessionStore(db *redis.Client) *SessionStore {
	return &SessionStore{db: db}
}

func (store *SessionStore) CreateSession(session *models.Session) error {
	// Marshal Data
	data, err := json.Marshal(session.Data)
    if err != nil {  return err  }

	// Update DB
	// key := "user_sessions:" + (string)(session.Data.UserID)
	err = store.db.Set("session:"+session.ID, data, time.Duration(session.ExpiryInNS)).Err()
	if err != nil {  return err  }

	return nil
}


func (store *SessionStore) DeleteSession(sessionID string) error {
	return store.db.Del(sessionID).Err()
}


func (store *SessionStore) GetSessionData(sessionID string) (*models.SessionData, error) {
	// Get Data as String
	val, err := store.db.Get("session:"+sessionID).Result()
    if err != nil { return nil, err }

	// Marshal Data To Struct
    var data models.SessionData
    err = json.Unmarshal([]byte(val), &data)
    if err != nil { return nil, err }

	return &data, nil
}