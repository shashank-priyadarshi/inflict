-- Amount Read
-- name: GetAmount :one
SELECT * FROM Amounts
WHERE id = $1 AND deleted = FALSE;

-- name: ListAmounts :many
SELECT * FROM Amounts
WHERE deleted = FALSE
ORDER BY created_at DESC;
