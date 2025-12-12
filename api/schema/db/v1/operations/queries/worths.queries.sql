-- Worths Read
-- name: GetWorth :one
SELECT * FROM Worths
WHERE id = $1 AND deleted = FALSE;

-- name: ListWorths :many
SELECT * FROM Worths
WHERE deleted = FALSE
ORDER BY created_at DESC;
