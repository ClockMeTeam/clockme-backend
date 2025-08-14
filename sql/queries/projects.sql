-- name: CreateProject :one
INSERT INTO projects(id, name)
VALUES(
       $1,
       $2
      )
RETURNING *;

-- name: GetAllProjects :many
SELECT * FROM projects;

-- name: GetProject :one
SELECT * FROM projects
WHERE projects.id = $1;

-- name: UpdateProject :one
UPDATE projects
SET
    updated_at = now(),
    name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE projects.id = $1;

-- name: DeleteAllProjects :exec
DELETE FROM projects;