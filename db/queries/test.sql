-- name: ListUsers :many
SELECT *
FROM users
ORDER BY email
LIMIT ?;