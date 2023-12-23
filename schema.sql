CREATE TABLE users (
	id TEXT PRIMARY KEY,
	account_name TEXT NOT NULL,
	password_hash TEXT NOT NULL
);

CREATE TABLE sessions (
	id TEXT PRIMARY KEY,
	user_id TEXT NOT NULL
);
