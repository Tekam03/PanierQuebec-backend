-- name: GetStoreMerchants :many
SELECT * FROM store_merchants;

-- name: GetStoreMerchantByID :one
SELECT * FROM store_merchants WHERE id = $1;

-- name: CreateStoreMerchant :one
INSERT INTO store_merchants (name, url) VALUES ($1, $2) RETURNING *;

-- name: UpdateStoreMerchant :one
UPDATE store_merchants
SET name = COALESCE($1, name),
    url = COALESCE($2, url)
WHERE id = $3
RETURNING *;

-- name: DeleteStoreMerchant :one
DELETE FROM store_merchants
WHERE id = $1
RETURNING *;

