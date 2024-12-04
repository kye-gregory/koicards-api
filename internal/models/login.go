package models

type Login struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func NewLogin(email, username, password string) *Login {
	// Create Struct
	login := Login{
		Email:    email,
		Username: username,
		Password: password,
	}

	// Return
	return &login
}