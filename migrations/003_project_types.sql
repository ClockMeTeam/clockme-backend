-- +goose Up
CREATE TABLE project_types(
    id UUID UNIQUE PRIMARY KEY NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS project_types;