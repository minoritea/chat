-- name: CreateUser :one
INSERT INTO users (id, account_name, password_hash) VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserByAccountName :one
SELECT * FROM users WHERE account_name = ?;

-- name: CreateSession :one
INSERT INTO sessions (id, user_id) VALUES (?, ?)
RETURNING *;

-- name: GetUserBySessionID :one
SELECT users.*
FROM users JOIN sessions ON sessions.user_id = users.id
WHERE sessions.id = ?;
