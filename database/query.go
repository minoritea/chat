package database

import "database/sql"

type Querier struct {
	db *sql.DB
}

type Session struct {
	UserID    string
	SessionID string
}

func (q Querier) CreateUser(accountName, passwordHash string) error {
	const query = "INSERT INTO users (id, account_name, password_hash) VALUES (?, ?, ?)"
	_, err := q.db.Exec(query, NewID(), accountName, passwordHash)
	return err
}

type User struct {
	ID           string
	AccountName  string
	PasswordHash string
}

func (q Querier) GetUserByAccountName(accountName string) (*User, error) {
	const query = "SELECT id, account_name, password_hash FROM users WHERE account_name = ?"
	var user User
	err := q.db.QueryRow(query, accountName).Scan(&user.ID, &user.AccountName, &user.PasswordHash)
	return &user, err
}

func (q Querier) CreateSession(userID string) (sessionID string, err error) {
	const query = "INSERT INTO sessions (id, user_id) VALUES (?, ?)"
	sessionID = NewID()
	_, err = q.db.Exec(query, sessionID, userID)
	return sessionID, err
}

func (q Querier) GetUserBySessionID(sessionID string) (*User, error) {
	const query = `
		SELECT
			users.id, users.account_name, users.password_hash
		FROM
			users
		INNER JOIN
			sessions
		ON
			users.id = sessions.user_id
		WHERE
			sessions.id = ?`
	var user User
	err := q.db.QueryRow(query, sessionID).Scan(&user.ID, &user.AccountName, &user.PasswordHash)
	return &user, err
}
