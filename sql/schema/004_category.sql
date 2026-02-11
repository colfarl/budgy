-- +goose Up
CREATE TABLE categories (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	is_income BOOLEAN NOT NULL,

	created_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	updated_at INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	deleted_at INTEGER,	
	deleted_reason TEXT
);
 
-- +goose Down
DROP TABLE categories;
