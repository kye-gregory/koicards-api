package postgres

import (
	"context"

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
	// Run Query & Scan Data
	var exists bool
	err := store.db.QueryRow(context.Background(), qCheckUsernameExists, username).Scan(&exists)

	// Check Errors & Return
	if err != nil { return false, err }
	return exists, nil
}


func (store *UserStore) IsEmailRegistered(email string) (bool, error) {
	// Run Query & Scan Data
	var exists bool
	err := store.db.QueryRow(context.Background(), qCheckEmailExists, email).Scan(&exists)

	// Check Errors & Return
	if err != nil { return false, err }
	return exists, nil
}


func (store *UserStore) CreateUser(u *models.User) error {
	// Run Query
	_, err := store.db.Exec(context.Background(), qInsertNewUser, u.Email, u.Username, u.Password)

	// Check Errors & Return
	if err != nil { return err }
	return nil
}


func (store *UserStore) VerifyEmail(email string) error {
	// Run Query
	_, err := store.db.Exec(context.Background(), qVerifyUserEmail, email)

	// Check Errors & Return
	if err != nil { return err }
	return nil
}

func (store *UserStore) getUser (query string, selector string) (*models.User, error) {
	// Run Query & Scan Data
	var user models.User
	var email, username, password string
	err := store.db.QueryRow(context.Background(), query, selector).Scan(
		&user.ID,
		&email,
		&username,
		&password,
		&user.IsVerified,
		&user.CreatedAt,
		&user.Status,
	)

	// Assign Struct Values
	user.Email = *userVO.NewEmailFromDB(email)
	user.Username = *userVO.NewUsernameFromDB(username)
	user.Password = *userVO.NewPasswordFromDB(password)

	// Check Errors & Return
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
	if err != nil { return nil, err }
	defer rows.Close()

	// Create User Slice
	var users []*models.User
	for rows.Next() {

		// Scan Data
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

		// Check Error
		if err != nil { return nil, err }

		// Assign Struct Values
		user.Email = *userVO.NewEmailFromDB(email)
		user.Username = *userVO.NewUsernameFromDB(username)
		user.Password = *userVO.NewPasswordFromDB(password)
		users = append(users, &user)
	}

	// Check Errors & Return
	if rows.Err() != nil { return nil, rows.Err() }
	return users, nil
}