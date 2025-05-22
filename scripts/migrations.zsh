#!/usr/bin/zsh
# migrations.zsh run database migrations with goose

# ZSH error handling
setopt ERR_EXIT
setopt PIPE_FAIL

if [[ -f .env ]]; then
  source .env
else
  echo ".env not found"
  exit 1
fi

if [[ ! -d "./migrations" ]]; then
  echo "migrations dir not found"
  exit 1
fi

# const
COMMAND=${1:-up}
DB_CONNECTION="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${HOST_DB_PORT}/${POSTGRES_DB}?sslmode=disable"

# Run goose
goose -dir ./migrations postgres "${DB_CONNECTION}" "${@}"

# verify
if [[ $? -eq 0 ]]; then
  echo "Migration '${COMMAND}' completed successfully!"
else
  echo "Migration '${COMMAND}' failed."
  exit 1
fi

