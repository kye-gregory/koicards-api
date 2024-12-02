package models

import (
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type User struct {
	ID			userVO.ID		`json:"id"`
	Email 		userVO.Email	`json:"email"`
	Username 	userVO.Username	`json:"username"`
	Password 	userVO.Password	`json:"-"`
	IsVerified	bool			`json:"isVerified"`
}

func NewUser(email userVO.Email, username userVO.Username, password userVO.Password) *User{
	user := User {
		Email: 		email,
		Username: 	username,
		Password: 	password,
		IsVerified: false,
	}

	// Return
	return &user
}