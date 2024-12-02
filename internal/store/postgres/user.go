package postgres

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kye-gregory/koicards-api/internal/models"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}

func (store *UserStore) IsUsernameRegistered(email string) (bool, error) {return false, fmt.Errorf("Not Implemented")}
func (store *UserStore) IsEmailRegistered(email string) (bool, error) {return false, fmt.Errorf("Not Implemented")}
func (store *UserStore) CreateUser(user *models.User) error  {return fmt.Errorf("Not Implemented")}
func (store *UserStore) ActivateUser(email string) error  {return fmt.Errorf("Not Implemented")}
func (store *UserStore) GetUser(identifier string) (*models.User, error)  {return nil, fmt.Errorf("Not Implemented")}
func (store *UserStore) GetAllUsers() ([]*models.User, error)  {return nil, fmt.Errorf("Not Implemented")}