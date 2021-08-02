-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1 LIMIT 1;

-- name: GetUserTags :one
SELECT tags FROM users WHERE id = $1; 

-- name: AddUserTag :exec
UPDATE users SET tags = tags || @tag::TEXT WHERE id = @id;

-- name: RemoveUserTag :exec
UPDATE users SET tags =  array_remove(tags, @tag::TEXT);

-- name: CreateUser :one
INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id;

-- name: GetUserPassword :one
SELECT password FROM users WHERE name = $1;
