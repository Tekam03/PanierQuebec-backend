-- name: GetStoreMerchants :many
SELECT * FROM store_merchants;

-- name: GetStoreMerchantByID :one
SELECT * FROM store_merchants WHERE id = $1;

-- name: CreateStoreMerchant :one
INSERT INTO store_merchants (name, url) VALUES ($1, $2) RETURNING *;

-- name: PatchStoreMerchant :one
UPDATE store_merchants
SET name = COALESCE(sqlc.narg('name'), name),
    url  = COALESCE(sqlc.narg('url'), url)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteStoreMerchant :one
DELETE FROM store_merchants
WHERE id = $1
RETURNING *;

