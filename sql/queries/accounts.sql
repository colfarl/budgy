-- name: CreateAccount :one
INSERT INTO accounts (name, user_id) 
VALUES (?, ?) 
RETURNING *;

-- name: GetAccountByName :one
SELECT * 
FROM accounts
WHERE name = ? and user_id = ?;

-- name: GetAllAccounts :many
SELECT * 
FROM accounts
WHERE user_id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE name = ? and user_id = ?;

