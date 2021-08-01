-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserTags :one
SELECT tags FROM users WHERE id = $1; 

-- name: UpdateUserTags :exec
UPDATE users SET tags = $1;