-- name: GetStoreByID :one
SELECT * FROM stores WHERE id = $1;

-- name: CreateStore :one
INSERT INTO stores (merchant_id, name, location) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateStore :one
UPDATE stores
SET name = COALESCE($2, name),
    location = COALESCE($3, location)
WHERE id = $1
RETURNING *;

-- name: DeleteStore :one
DELETE FROM stores
WHERE id = $1
RETURNING *;
