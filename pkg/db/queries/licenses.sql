-- name: InsertLicense :one
INSERT INTO "licenses" (product, user, credentials)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetLicensesByProductAndUser :many
SELECT * FROM "licenses"
WHERE product = ? AND user = ?
ORDER BY id DESC;