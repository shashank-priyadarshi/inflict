-- Maintainance CRUD
-- name: CreateMaintainance :one
INSERT INTO Maintainances (id, wealth_id, type, name, cost_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateMaintainance :one
UPDATE Maintainances
SET wealth_id = $2,
    type = $3,
    name = $4,
    cost_id = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted = FALSE
RETURNING *;

-- name: DeleteMaintainance :exec
UPDATE Maintainances
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
