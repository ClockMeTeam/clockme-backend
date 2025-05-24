-- +goose Up
CREATE TABLE workspaces(
      id TEXT UNIQUE PRIMARY KEY,
      name TEXT UNIQUE NOT NULL ,
      email TEXT UNIQUE NOT NULL,
      created_at TIMESTAMP NOT NULL,
      update_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTs workspaces;