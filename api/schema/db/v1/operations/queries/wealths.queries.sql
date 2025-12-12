-- Wealths Read
-- name: GetWealth :one
SELECT * FROM Wealths
WHERE id = $1 AND deleted = FALSE;

-- name: GetWealthsByWorth :many
SELECT * FROM Wealths
WHERE worth_id = $1 AND deleted = FALSE
ORDER BY created_at DESC;

-- name: GetWealthsByType :many
SELECT * FROM Wealths
WHERE type = $1 AND deleted = FALSE
ORDER BY created_at DESC;

-- name: ListWealths :many
SELECT * FROM Wealths
WHERE deleted = FALSE
ORDER BY created_at DESC;
