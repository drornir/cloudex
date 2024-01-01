-- name: InsertLicense :one
INSERT INTO "licenses" (product, user, credentials)
VALUES (?, ?, ?)
RETURNING *;