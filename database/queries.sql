-- name: CreateUser :one
INSERT INTO users (id, account) VALUES (?, ?)
RETURNING *;

-- name: GetUserByAccount :one
SELECT * FROM users WHERE account = ?;

-- name: CreateSession :one
INSERT INTO sessions (id, user_id, expires_at) VALUES (?, ?, ?)
RETURNING *;

-- name: GetUserBySessionID :one
SELECT users.*
FROM users JOIN sessions ON sessions.user_id = users.id
WHERE sessions.id = ? and sessions.expires_at > current_timestamp;

-- name: CreateMessage :one
INSERT INTO messages (id, user_id, message, created_at) VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListNewestMessages :many
SELECT messages.id, messages.message, messages.created_at, users.account
FROM messages JOIN users ON messages.user_id = users.id
ORDER BY messages.id DESC LIMIT ?;

-- name: ListOldestMessages :many
SELECT messages.id, messages.message, messages.created_at, users.account
FROM messages JOIN users ON messages.user_id = users.id
ORDER BY messages.id ASC LIMIT ?;

-- name: ListMessagesBeforeID :many
SELECT messages.id, messages.message, messages.created_at, users.account
FROM messages JOIN users ON messages.user_id = users.id
WHERE messages.id < ?
ORDER BY messages.id DESC LIMIT ?;

-- name: ListMessagesAfterID :many
SELECT messages.id, messages.message, messages.created_at, users.account
FROM messages JOIN users ON messages.user_id = users.id
WHERE messages.id > ?
ORDER BY messages.id ASC LIMIT ?;
