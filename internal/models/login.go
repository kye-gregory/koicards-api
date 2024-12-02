package models

type Login struct {
	Identifier string `json:"identifier"`
	Password   string `json:"-"`
}

func NewLogin(email string, username string, password string) *Login {
	// Determine Identifier
	identifier := email
	if username != "" {
		identifier = username
	}

	// Create Struct
	login := Login{
		Identifier: identifier,
		Password:   password,
	}

	// Return
	return &login
}