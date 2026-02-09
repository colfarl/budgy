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
	WHERE u.name = ? AND a.name = ?
	), ?, ?, ?, ?) 
RETURNING *;

-- name: GetAccountTxnFromNames :many
SELECT t.*
FROM users u JOIN accounts a ON u.ID = a.user_id
			 JOIN transactions t ON t.account_id = a.id
WHERE u.name = ? AND a.name = ?;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = ?;
