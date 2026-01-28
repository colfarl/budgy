-- +goose Up
CREATE TABLE session (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	username TEXT DEFAULT NULL,
	last_opened INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),

	FOREIGN KEY (username) REFERENCES users(name) ON DELETE SET NULL
);

INSERT INTO session (id) VALUES (1); -- Singleton row pattern

-- +goose Down
DROP TABLE session;
