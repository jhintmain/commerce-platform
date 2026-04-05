-- name: InsertBrand :execlastid
INSERT INTO brands (name)
VALUES (?);

-- name: GetBrand :one
SELECT *
FROM brands
WHERE id = ?;

-- name: GetBrandByName :one
SELECT *
FROM brands
WHERE name = ?;

-- name: UpdateBrand :exec
UPDATE brands
SET name = ?
WHERE id = ?;

-- name: DeleteBrand :exec
UPDATE brands
SET deleted_at = ?
WHERE id = ?;