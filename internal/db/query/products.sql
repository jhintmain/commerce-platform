-- name: CreateProduct :execlastid
INSERT INTO products (brand_id,
                      name,
                      description,
                      base_price,
                      sale_status)
VALUES (?, ?, ?, ?, ?);

-- name: GetProduct :one
SELECT * FROM products
WHERE id = ?;

-- name: GetProducts :many
SELECT * FROM products
WHERE deleted_at IS NULL;

-- name: GetProductsByBrand :many
SELECT * FROM products
WHERE brand_id = ? AND deleted_at IS NULL;

-- name: UpdateProduct :exec
UPDATE products
SET name        = ?,
    description = ?,
    base_price  = ?,
    sale_status = ?
WHERE id = ?;

-- name: DeleteProduct :exec
UPDATE products
SET deleted_at = ?
WHERE id = ?;