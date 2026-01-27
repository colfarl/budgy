-- +goose Up
CREATE TABLE session (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	username TEXT,
	last_opened INTEGER NOT NULL,

	FOREIGN KEY (username) REFERENCES users(name) ON DELETE SET NULL
);
INSERT OR IGNORE INTO session (id) VALUES (1); -- Singleton row pattern

-- +goose Down
DROP TABLE session;
