-- +goose Up
CREATE TABLE category_transaction (
	transaction_id INTEGER NOT NULL UNIQUE,
	category_id INTEGER NOT NULL,

	FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
	FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE category_transaction;
