package postgres

const (
	qCheckEmailExists = `
	SELECT EXISTS (
        SELECT 1
        FROM user_account
        WHERE username = $1
    )`

	qCheckUsernameExists = `
	SELECT EXISTS (
        SELECT 1
        FROM user_account
        WHERE username = $1
    )`

	qInsertNewUser = `
	INSERT INTO user_account (email, username, password_hash) 
	VALUES ($1, $2, $3);`

	qVerifyUserEmail = `
	UPDATE user_account
	SET is_email_verified = 1
	WHERE email = $1`

	qGetUserByEmail = `
	SELECT id, email, username, password_hash, is_email_verified, created_at, account_status
	FROM user_account
	WHERE email = $1`

	qGetUserByUsername = `
	SELECT id, email, username, password_hash, is_email_verified, created_at, account_status
	FROM user_account
	WHERE email = $1`

	qGetAllUsers = `
	SELECT id, email, username, password_hash, is_email_verified, created_at, account_status
	FROM user_account`
)