#!/usr/bin/zsh

# Build the go application
go build -o ftfclokify cmd/api/main.go

# Run docker compose
docker compose up -d --build
