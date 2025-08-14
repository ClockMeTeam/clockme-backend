-- name: AddUserToProject :one
INSERT INTO projects_users (user_id, project_id)
VALUES (
        $1,
        $2
       )
RETURNING *;

-- name: RemoveUserFromProject :exec
DELETE FROM projects_users
WHERE user_id = $1 AND project_id = $2;

-- name: ListUsersForProject :many
SELECT users.* FROM users
JOIN projects_users ON users.id = projects_users.user_id
WHERE projects_users.project_id = $1;

-- name: ListProjectsForUser :many
SELECT projects.* FROM projects
JOIN projects_users ON projects.id = projects_users.project_id
WHERE projects_users.user_id = $1;

