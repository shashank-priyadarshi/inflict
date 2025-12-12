-- name: GetMember :one
SELECT * FROM Members
WHERE id = $1 LIMIT 1;

-- name: ListMembers :many
SELECT * FROM Members
ORDER BY name;
