-- name: GetExternalProducts :many
SELECT
    ep.id,
    ep.source,
    ep.external_id,
    ep.name,
    ep.description,
    ep.brand,
    ep.matched_store_product_id,
    ep.scraped_at
FROM external_products ep;

-- name: GetExternalProductByExternalID :one
SELECT
    ep.id,
    ep.source,
    ep.external_id,
    ep.name,
    ep.description,
    ep.brand,
    ep.matched_store_product_id,
    ep.scraped_at
FROM external_products ep
WHERE ep.external_id = $1;

-- name: GetExternalProductsDetailed :many
SELECT
    ep.id AS ep_id,
    ep.source AS ep_source,
    ep.external_id AS ep_external_id,
    ep.name AS ep_name,
    ep.description AS ep_description,
    ep.brand AS ep_brand,
    ep.matched_store_product_id AS ep_matched_store_product_id,
    ep.scraped_at AS ep_scraped_at,
    sp.id AS sp_id,
    sp.store_id AS sp_store_id,
    sp.specific_products_id AS sp_specific_products_id,
    sp.price AS sp_price,
    sp.last_updated AS sp_last_updated
FROM external_products ep
LEFT JOIN store_products sp ON ep.matched_store_product_id = sp.id;

-- name: CreateExternalProduct :one
INSERT INTO external_products (
    source,
    external_id,
    name,
    description,
    brand,
    matched_store_product_id
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: DeleteExternalProduct :one
DELETE FROM external_products
WHERE id = $1
RETURNING *;
