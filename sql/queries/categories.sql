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

-- name: TxnsByCategory :many
SELECT c.name, SUM(t.amount) AS total
FROM users u 
	JOIN accounts a ON a.user_id = u.id 
	JOIN transactions t ON t.account_id = a.id
	JOIN category_transaction ct ON ct.transaction_id = t.id
	JOIN categories c on ct.category_id = c.id 
WHERE u.name = ? 
	AND t.occurred_at >= ? 
	AND t.occurred_at <= ?
	AND c.is_income = ?
GROUP BY c.name;
				
