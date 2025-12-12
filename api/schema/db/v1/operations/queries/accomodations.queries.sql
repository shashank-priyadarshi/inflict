-- name: ListAccomodations :many
SELECT * FROM Accomodations
WHERE member_id = @member_id AND deleted = FALSE;
