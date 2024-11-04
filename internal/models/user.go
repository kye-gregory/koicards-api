package models

// Using Placeholder Structs As Mock DB
type Login struct {
	HashedPassword string
	SessionToken string
	CSRFToken string
}