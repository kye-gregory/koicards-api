package models

// Using Placeholder Structs As Mock DB
type Login struct {
	User 			User
	SessionToken 	string
	CSRFToken 		string
}

type User struct {
	ID			int
	Username 	string
	Password 	string
	Email 		string
}