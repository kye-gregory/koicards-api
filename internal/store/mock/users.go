package mock

import (
	"maps"
	"slices"

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

func (store *UserStore) CreateUser(u *models.User) error {
	id := len(store.users)
	u.ID = id
	store.users[id] = u
	return nil
}


func (store *UserStore) ActivateUser(email string) error {
	user, err := store.GetUser(email)
	if (err != nil) { return err }

	user.IsVerified = true
	return nil
}

func (store *UserStore) GetUser(identifier string) (*models.User, error) {
	for _, user := range store.users {
		hasEmail := (user.Email.String() == identifier)
		hasUsername := (user.Username.String() == identifier)
		if (!hasEmail && !hasUsername) { continue }
		return user, nil
	}

	return nil, nil
}

func (store *UserStore) GetAllUsers() ([]*models.User, error) {
	return slices.Collect(maps.Values(store.users)), nil
}