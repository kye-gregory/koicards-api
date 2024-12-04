package mock

import (
	"fmt"

	"github.com/kye-gregory/koicards-api/internal/models"
)

type SessionStore struct {
	sessions []models.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{sessions: make([]models.Session, 0)}
}

func (store *SessionStore) CreateSession(session *models.Session) error {
	store.sessions = append(store.sessions, *session)
	return nil
}

func (store *SessionStore) DeleteSession(sessionID string) error {
	for i, session := range store.sessions {
		if (session.ID != sessionID) { continue }
		store.sessions = append(store.sessions[:i], store.sessions[i+1:]...)
		return nil
	}

	return fmt.Errorf("session id not found")
}


func (store *SessionStore) GetSessionData(sessionID string) (*models.SessionData, error) {
	for _, session := range store.sessions {
		if (session.ID != sessionID) { continue }
		return &session.Data, nil
	}
	return nil, nil
}