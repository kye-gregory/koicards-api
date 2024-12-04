package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kye-gregory/koicards-api/internal/models"
	userVO "github.com/kye-gregory/koicards-api/internal/valueobjects/user"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}


func (store *UserStore) IsUsernameRegistered(username string) (bool, error) {
	var exists bool
	err := store.db.QueryRow(context.Background(), qCheckUsernameExists, username).Scan(&exists)
	if err != nil { return false, err }

	return exists, nil
}


func (store *UserStore) IsEmailRegistered(email string) (bool, error) {
	var exists bool
	err := store.db.QueryRow(context.Background(), qCheckEmailExists, email).Scan(&exists)
	if err != nil { return false, err }

	return exists, nil
}


func (store *UserStore) CreateUser(u *models.User) error {
	_, err := store.db.Exec(context.Background(), qInsertNewUser, u.Email, u.Username, u.Password)
	if err != nil { return err }

	return nil
}


func (store *UserStore) VerifyEmail(email string) error {
	_, err := store.db.Exec(context.Background(), qVerifyUserEmail, email)
	if err != nil { return err }

	return nil
}

func (store *UserStore) getUser (query string, selector string) (*models.User, error) {
	var user models.User
	err := store.db.QueryRow(context.Background(), query, selector).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.IsVerified,
		&user.CreatedAt,
		&user.Status,
	)

	if err != nil { return nil, err }
	return &user, nil
}

func (store *UserStore) GetUserByEmail(email string) (*models.User, error) {
	return store.getUser(qGetUserByEmail, email)
}


func (store *UserStore) GetUserByUsername(username string) (*models.User, error) {
	return store.getUser(qGetUserByUsername, username)
}


func (store *UserStore) GetAllUsers() ([]*models.User, error) {
	// Run Query
	rows, err := store.db.Query(context.Background(), qGetAllUsers)
	if err != nil { return nil, fmt.Errorf("failed to query users: %w", err) }
	defer rows.Close()

	// Create Slice
	var users []*models.User
	for rows.Next() {
		var user models.User
		var email, username, password string
		err := rows.Scan(
			&user.ID,
			&email,
			&username,
			&password,
			&user.IsVerified,
			&user.CreatedAt,
			&user.Status,
		)
		if err != nil { return nil, fmt.Errorf("failed to scan user: %w", err) }
		user.Email = *userVO.NewEmailFromDB(email)
		user.Username = *userVO.NewUsernameFromDB(username)
		user.Password = *userVO.NewPasswordFromDB(password)
		users = append(users, &user)
	}

	// Final Error Check
	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating rows: %w", rows.Err())
	}

	return users, nil
}