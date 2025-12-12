-- Returns CRUD
-- name: CreateReturn :one
INSERT INTO Returns (id, wealth_id, name, rate_type, rate_value, duration, maturity_corpus_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateReturn :one
UPDATE Returns
SET wealth_id = $2,
    name = $3,
    rate_type = $4,
    rate_value = $5,
    duration = $6,
    maturity_corpus_id = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted = FALSE
RETURNING *;

-- name: DeleteReturn :exec
UPDATE Returns
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
