-- +goose Up
CREATE TABLE projects(
    id UUID UNIQUE PRIMARY KEY NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    name TEXT NOT NULL,
    type_id UUID REFERENCES project_types(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS projects;