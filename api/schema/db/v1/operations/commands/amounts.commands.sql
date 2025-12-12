-- Amount CRUD
-- name: CreateAmount :one
INSERT INTO Amounts (id, type, name, sender, receiver, value, currency)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateAmount :one
UPDATE Amounts
SET type = $2,
    name = $3,
    sender = $4,
    receiver = $5,
    value = $6,
    currency = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted = FALSE
RETURNING *;

-- name: DeleteAmount :exec
UPDATE Amounts
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
