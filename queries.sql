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

-- name: CreateMessage :one
INSERT INTO messages (id, user_id, message, created_at) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListMessages :many
SELECT messages.id, messages.user_id, messages.message, messages.created_at, users.account_name
FROM messages JOIN users ON messages.user_id = users.id
ORDER BY created_at DESC LIMIT ?;
