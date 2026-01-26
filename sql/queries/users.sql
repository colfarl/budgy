-- name: CreateUser :one
INSERT INTO users (name) 
VALUES (?) 
RETURNING *;

-- name: GetUserByName :one
SELECT * 
FROM users 
WHERE name = ?;

-- name: GetAllUserNames :many
SELECT name 
FROM users;

-- name: DeleteAccountByName :exec
DELETE FROM users 
WHERE name = ?;

