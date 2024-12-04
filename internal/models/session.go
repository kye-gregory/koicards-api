package models

import (
	"time"

	"github.com/google/uuid"
)
type Session struct {
	ID   string
	Data SessionData
	ExpiryInNS int64
}

type SessionData struct {
	UserID int
	CSRFToken string
}

func NewSession(data SessionData) *Session {
	return &Session{
		ID: uuid.New().String(),
		Data: data,
		ExpiryInNS: time.Hour.Nanoseconds() * 24,
	}
}

func NewSessionData(userID int, csrfToken string) *SessionData {
	return &SessionData{
		UserID: userID,
		CSRFToken: csrfToken,
	}
}