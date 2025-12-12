-- Maintainance Read
-- name: GetMaintainance :one
SELECT * FROM Maintainances
WHERE id = $1 AND deleted = FALSE;

-- name: GetMaintainancesByWealth :many
SELECT * FROM Maintainances
WHERE wealth_id = $1 AND deleted = FALSE
ORDER BY created_at DESC;

-- name: ListMaintainances :many
SELECT * FROM Maintainances
WHERE deleted = FALSE
ORDER BY created_at DESC;
