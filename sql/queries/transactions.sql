-- name: CreateTransaction :one
INSERT INTO transactions (
	account_id,
	is_income,
	amount,
	description,
	occurred_at
) VALUES (?, ?, ?, ?, ?) 
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = ?;
