package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID   string
	User int
	Expiry time.Time
}

func NewSession(userID int) *Session {
	return &Session{
		ID: uuid.New().String(),
		User: userID,
		Expiry: time.Now().Add(time.Hour * 24),
	}
}