package server

// Using Placeholder Structs As Mock DB
type Login struct {
	HashedPassword string
	SessionToken string
	CSRFToken string
}

var users = map[string]Login{}