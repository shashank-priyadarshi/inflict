-- Worths CRUD
-- name: CreateWorth :one
INSERT INTO Worths (id)
VALUES ($1)
RETURNING *;

-- name: UpdateWorth :one
UPDATE Worths
SET updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted = FALSE
RETURNING *;

-- name: DeleteWorth :exec
UPDATE Worths
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
