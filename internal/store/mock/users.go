package mock

import "github.com/kye-gregory/koicards-api/internal/models"

type UserStore struct {
	users map[int]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[int]*models.User)}
}

func (repo *UserStore) UserExists(email string) (bool, error) {
	return false, nil
}

func (repo *UserStore) CreateUser(user *models.User) error {
	repo.users[user.ID] = user
	return nil
}