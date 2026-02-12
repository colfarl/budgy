-- name: CreateCategory :one
INSERT INTO categories (name, is_income) 
VALUES (?, ?) 
RETURNING *;

-- name: GetAllCategories :many
SELECT * 
FROM categories;

-- name: GetCategoriesByType :many
SELECT * 
FROM categories
WHERE is_income = ?;

-- name: DeleteCategory :exec
DELETE FROM categories 
WHERE id = ?;

