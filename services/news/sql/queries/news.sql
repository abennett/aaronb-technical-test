-- name: ListNewsByTags :many
SELECT * FROM news WHERE tags @> @tags::text[] ORDER BY "timestamp" DESC;

-- name: ListNewByTagsPaged :many
SELECT * FROM news WHERE "timestamp" < @timestamp AND tags @> @tags::text[] ORDER BY "timestamp" DESC;

-- name: CreateNews :one
INSERT INTO news (title, "timestamp", tags) VALUES ($1, $2, $3) RETURNING id;
