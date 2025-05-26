-- +goose Up
CREATE TABLE users(
    id UUID UNIQUE PRIMARY KEY NOT NULL DEFAULT pg_catalog.gen_random_uuid(),
    clockify_id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTs users;