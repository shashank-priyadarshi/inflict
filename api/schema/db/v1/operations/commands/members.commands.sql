-- name: CreateMember :one
INSERT INTO Members (
  id, name, type, net_worth_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateMember :one
UPDATE Members
SET name = $2,
    type = $3,
    net_worth_id = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteMember :exec
UPDATE Members
SET deleted = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;