-- name: SaveUser :exec
INSERT INTO users (email)
VALUES (?);