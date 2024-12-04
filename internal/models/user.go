package models

import (
	"time"

	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type User struct {
	ID			int				`json:"id"`
	Email 		userVO.Email	`json:"email"`
	Username 	userVO.Username	`json:"username"`
	Password 	userVO.Password	`json:"-"`
	IsVerified	bool			`json:"isVerified"`
	CreatedAt	time.Time		`json:"createdAt"`
	Status		string			`json:"status"`
}

func NewUser(email userVO.Email, username userVO.Username, password userVO.Password) *User{
	user := User {
		ID:			-1,
		Email: 		email,
		Username: 	username,
		Password: 	password,
		IsVerified: false,
	}

	// Return
	return &user
}