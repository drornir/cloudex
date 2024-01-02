-- name: GetUser :one
SELECT * FROM "users"
WHERE id = ? LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM "users";

-- name: InsertUser :one
INSERT INTO "users" (email)
VALUES (?)
RETURNING *;