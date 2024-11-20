package models

import (
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type User struct {
	ID			int
	Email 		userVO.Email
	Username 	userVO.Username
	Password 	userVO.Password
	IsVerified	bool
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