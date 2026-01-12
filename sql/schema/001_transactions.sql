-- +goose Up
CREATE TABLE transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT
);

-- +goose Down
DROP TABLE transactions;
