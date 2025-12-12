-- name: CreateAccomodation :one
INSERT INTO Accomodations (
    id, member_id, type, address, cost_id
) VALUES (
    @id, @member_id, @type, @address, @cost_id
) RETURNING *;

-- name: UpdateAccomodation :one
UPDATE Accomodations
SET
    type = COALESCE(sqlc.narg('type'), type),
    address = COALESCE(sqlc.narg('address'), address),
    cost_id = COALESCE(sqlc.narg('cost_id'), cost_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteAccomodation :exec
UPDATE Accomodations
SET deleted = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = @id;
