-- name: GetProjectByName :one
SELECT * FROM projects WHERE name = $1;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: UpdateProject :one
UPDATE projects
SET
    name = $1,
    type_id = COALESCE($2, type_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;

-- name: CreateProject :one
INSERT INTO projects(name, clockify_id) VALUES ($1, $2)
RETURNING *;

-- name: DeleteProjectByName :exec
DELETE FROM projects WHERE name = $1;

-- name: GetProjectType :one
SELECT pt.name FROM projects p
LEFT JOIN project_types pt  ON p.type_id = pt.id
WHERE p.id = $1;

-- name: GetProjectTypeByClockifyId :one
SELECT pt.* FROM projects p
LEFT JOIN project_types pt  ON p.type_id = pt.id
WHERE p.clockify_id = $1;

-- name: UpdateProjectType :one
UPDATE projects
SET
    type_id = $1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $2
RETURNING *;

-- name: GetProjectUsers :many
SELECT u.* FROM users u
JOIN user_projects up ON u.id = up.user_id
WHERE up.project_id = $1;
