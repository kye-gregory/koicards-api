package mock

import "github.com/kye-gregory/koicards-api/internal/models"

type UserStore struct {
	users map[int]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[int]*models.User)}
}

func (s *UserStore) UserExists(email string) (bool, error) {
	return false, nil
}

func (s *UserStore) CreateUser(user *models.User) error {
	id := len(s.users)
	user.ID = id
	s.users[id] = user
	return nil
}