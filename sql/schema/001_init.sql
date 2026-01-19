-- +goose Up
CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,

	created_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	updated_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	deleted_at TEXT,	
	deleted_reason TEXT
);

CREATE TABLE accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	name TEXT NOT NULL UNIQUE,

	created_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	updated_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	deleted_at TEXT,	
	deleted_reason TEXT,

	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,

	account_id INTEGER NOT NULL,
	is_income BOOLEAN NOT NULL DEFAULT FALSE,	
	amount INTEGER NOT NULL,
	description TEXT NOT NULL,
	occurred_at TEXT NOT NULL,

	created_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	updated_at TEXT NOT NULL DEFAULT (datetime('now', 'utc')),
	deleted_at TEXT,	
	deleted_reason TEXT,

	FOREIGN KEY(account_id) REFERENCES accounts(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE transactions;
DROP TABLE accounts;
