package db

import (
	"context"
	"fmt"
	"log"

	"github.com/R3iwan/blog-app/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(cfg *config.Config) (*pgxpool.Pool, error) {
	postgresCfg := cfg.Postgres

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		postgresCfg.User, postgresCfg.Password,
		postgresCfg.Host, postgresCfg.Port,
		postgresCfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database %w", err)
	}

	DB = pool

	log.Println("Database connected successfully")
	return pool, nil
}
