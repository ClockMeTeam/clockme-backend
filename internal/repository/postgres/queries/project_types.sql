-- name: GetProjectProjectType :one
SELECT * FROM project_types WHERE name = $1;

-- name: GetProjectTypes :many
SELECT * FROM project_types;

-- name: UpdateProjectProjectType :one
UPDATE project_types
SET
    name = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE name = $2
RETURNING *;

-- name: CreateProjectType :one
INSERT INTO project_types(name) VALUES ($1)
RETURNING *;

-- name: DeleteProjectTypeByName :exec
DELETE FROM project_types WHERE name = $1
RETURNING *;