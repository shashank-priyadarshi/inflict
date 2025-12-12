-- Returns Read
-- name: GetReturn :one
SELECT * FROM Returns
WHERE id = $1 AND deleted = FALSE;

-- name: GetReturnsByWealth :many
SELECT * FROM Returns
WHERE wealth_id = $1 AND deleted = FALSE
ORDER BY created_at DESC;

-- name: ListReturns :many
SELECT * FROM Returns
WHERE deleted = FALSE
ORDER BY created_at DESC;
