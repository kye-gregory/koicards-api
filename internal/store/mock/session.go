package mock

import (
	"fmt"

	"github.com/kye-gregory/koicards-api/internal/models"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type SessionStore struct {
	sessions []models.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make([]models.Session, 0)}
}

func (store *SessionStore) CreateSession(userID userVO.ID) (*models.Session, error) {
	session := models.NewSession(userID)
	store.sessions = append(store.sessions, *session)
	return session, nil
}

func (store *SessionStore) DeleteSession(sessionID string) error {
	for i, session := range store.sessions {
		if (session.ID != sessionID) { continue }
		store.sessions = append(store.sessions[:i], store.sessions[i+1:]...)
		return nil
	}

	return fmt.Errorf("session id not found")
}


func (store *SessionStore) VerifySession(sessionID string) (bool, error) {
	for _, session := range store.sessions {
		if (session.ID != sessionID) { continue }
		return true, nil
	}
	return false, nil
}