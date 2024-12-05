package models

import (
	"time"

	"github.com/google/uuid"
)
type Session struct {
	ID   		string		`json:"id"`
	Data 		SessionData	`json:"data"`
	ExpiryInNS 	int64		`json:"expiry_in_ns"`
}

type SessionData struct {
	UserID 	 	int		`json:"user_id"`
	CSRFToken	string	`json:"csrf_token"`
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