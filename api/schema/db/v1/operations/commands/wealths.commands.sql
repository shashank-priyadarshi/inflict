-- Wealths CRUD
-- name: CreateWealth :one
INSERT INTO Wealths (id, worth_id, type, name, value_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateWealth :one
UPDATE Wealths
SET worth_id = $2,
    type = $3,
    name = $4,
    value_id = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted = FALSE
RETURNING *;

-- name: DeleteWealth :exec
UPDATE Wealths
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
