-- +goose Up
CREATE TABLE user_projects(
    id  UUID UNIQUE PRIMARY KEY NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    projects_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, projects_id)
);

-- index for faster lookups
CREATE INDEX idx_user_projects_user_id ON user_projects(user_id)
CREATE INDEX idx_user_projects_projects_id ON user_projects(projects_id)

-- +goose Down
DROP TABLE IF EXISTS user_projects;