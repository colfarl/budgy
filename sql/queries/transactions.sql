-- name: CreateTransaction :one
INSERT INTO transactions (
	account_id,
	is_income,
	amount,
	description,
	occurred_at
) VALUES (?, ?, ?, ?, ?) 
RETURNING *;

-- name: CreateTransactionFromNames :one
INSERT INTO transactions (
	account_id,
	is_income,
	amount,
	description,
	occurred_at
) VALUES ((
	SElECT a.id 
	FROM accounts a 
		JOIN users u on a.user_id = u.ID
	WHERE u.name = @username AND a.name = @account_name
	), @is_income, @amount, @description, @occurred_at) 
RETURNING *;

-- name: GetAccountTxnFromNames :many
SELECT t.*
FROM users u JOIN accounts a ON u.ID = a.user_id
			 JOIN transactions t ON t.account_id = a.id
WHERE u.name = sqlc.arg(username) AND a.name = sqlc.arg(account_name);

-- name: GetAccountUncategorizedTxnFromNames :many
SELECT t.*
FROM users u JOIN accounts a ON u.ID = a.user_id
			 JOIN transactions t ON t.account_id = a.id
WHERE u.name = sqlc.arg(username) AND a.name = sqlc.arg(account_name)
	AND t.id NOT IN (SELECT transaction_id FROM category_transaction);

-- name: GetTxnFromID :one
SELECT *
FROM transactions
where id = ?;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = ?;

-- name: CheckIfTransactionCategorized :one
SELECT t.id
FROM transactions t 
WHERE t.id = ? AND t.id in (SELECT transaction_id FROM category_transaction);

-- name: UncategorizeTransaction :exec
DELETE FROM category_transaction
WHERE transaction_id = ?;

-- name: CategorizeTransactionByName :exec
INSERT INTO category_transaction (
	transaction_id, 
	category_id
) VALUES (
	?, (SELECT id FROM categories WHERE name = ?)
);

-- name: SumAccountTxnFromNames :one
SELECT
	SUM(
		CASE
			WHEN is_income = 1 THEN amount
			ELSE -amount
		END
	) as balance
FROM users u JOIN accounts a ON u.ID = a.user_id
			 JOIN transactions t ON t.account_id = a.id
WHERE u.name = sqlc.arg(username) AND a.name = sqlc.arg(account_name);

-- name: SumAccountFromUsername :many
SELECT
	a.name,
	SUM(
		CASE
			WHEN is_income = 1 THEN amount
			ELSE -amount
		END
	) as balance
FROM users u JOIN accounts a ON u.ID = a.user_id
			 JOIN transactions t ON t.account_id = a.id
WHERE u.name = sqlc.arg(username) 
GROUP BY a.name;

