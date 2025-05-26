-- name: CreateUser :one
INSERT INTO users (id, clockify_id, name, email, update_at)
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
       )
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;