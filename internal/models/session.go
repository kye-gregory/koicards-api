package models

import (
	"time"

	"github.com/google/uuid"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type Session struct {
	ID   string
	User userVO.ID
	Expiry time.Time
}

func NewSession(userID userVO.ID) *Session {
	return &Session{
		ID: uuid.New().String(),
		User: userID,
		Expiry: time.Now().Add(time.Hour * 24),
	}
}