-- +goose Up
CREATE TABLE projects_users(
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, project_id)
);

-- +goose down
DROP TABLE IF EXISTS projects_users;
