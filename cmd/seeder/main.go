package main

import (
	"context"
	"database/sql"
	"github.com/clockme/clockme-backend/internal/db"
	"github.com/clockme/clockme-backend/internal/logger"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	logger.Init()
	ctxBg := context.Background()

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal().Msg("DB_SOURCE is not set")
	}

	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	log.Info().Msg("Successfully connected to the database!")

	queries := db.New(conn)

	for i := 0; i < 3; i++ {
		newUser, err := queries.CreateUser(ctxBg, db.CreateUserParams{
			ID:    uuid.New(),
			Name:  faker.Name(),
			Email: faker.Email(),
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to create user")
		} else {
			log.Info().
				Stringer("id", newUser.ID).
				Str("name", newUser.Name).
				Str("email", newUser.Email).
				Msg("Created user")
		}
	}
}
