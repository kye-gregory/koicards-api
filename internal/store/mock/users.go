package mock

import (
	"github.com/kye-gregory/koicards-api/internal/models"
)

type UserStore struct {
	users map[int]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[int]*models.User)}
}

func (store *UserStore) IsEmailRegistered(email string) (bool, error) {
	for _, v := range store.users {
		if (v.Email.String() == email) {
			return true, nil
		}
	}
	return false, nil
}

func (store *UserStore) IsUsernameRegistered(username string) (bool, error) {
	for _, v := range store.users {
		if (v.Username.String() == username) {
			return true, nil
		}
	}
	return false, nil
}

func (store *UserStore) CreateUser(user *models.User) error {
	id := len(store.users)
	user.ID = id
	store.users[id] = user
	return nil
}
