-- name: GetStores :many
SELECT * FROM stores;

-- name: GetStoresDetailed :many
SELECT s.*, m.name AS merchant_name, m.url AS merchant_url
FROM stores s
LEFT JOIN store_merchants m ON s.merchant_id = m.id;

-- name: GetStoreByID :one
SELECT * FROM stores WHERE id = sqlc.arg('id');

-- name: GetStoreByIDDetailed :one
SELECT s.*, m.name AS merchant_name, m.url AS merchant_url
FROM stores s
LEFT JOIN store_merchants m ON s.merchant_id = m.id
WHERE s.id = sqlc.arg('id');

-- name: CreateStore :one
INSERT INTO stores (merchant_id, name, location) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateStore :one
UPDATE stores
SET name     = COALESCE(sqlc.narg('name'), name),
    location = COALESCE(sqlc.narg('location'), location)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteStore :one
DELETE FROM stores
WHERE id = sqlc.arg('id')
RETURNING *;
