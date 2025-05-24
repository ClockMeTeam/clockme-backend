package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maevlava/ftf-clockify/internal/config"
	"log"
)

func NewConnectionPool(dbCfg config.DatabaseConfig) (pool *pgxpool.Pool, err error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s&statement_cache_mode=describe",
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
		dbCfg.SSLMode,
	)
	log.Printf("Connecting to database: \n%s", dsn)
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors.New("error parsing connection string")
	}

	// poolConfig.MaxConns = 10
	// poolConfig.MinConns = 2
	// poolConfig.MaxConnLifetime = time.Hour
	// poolConfig.MaxConnIdleTime = time.Minute * 30
	// poolConfig.HealthCheckPeriod = time.Minute

	dbPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.New("error creating connection pool")
	}

	// verify
	if err := dbPool.Ping(context.Background()); err != nil {
		dbPool.Close()
		return nil, errors.New("error pinging database")
	}

	return dbPool, nil
}
