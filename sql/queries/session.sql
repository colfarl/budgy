-- name: LoadSession :one
SELECT username
FROM session;

-- name: LoginSession :exec
UPDATE session
SET username = ?1, last_opened = (strftime('%s', 'now'))
WHERE id = 1;

-- name: UpdateSessionLastOpened :exec
UPDATE session
SET last_opened = (strftime('%s', 'now'))
WHERE id = 1;


-- name: LogoutSession :exec
UPDATE session
SET username = NULL
WHERE id = 1;
