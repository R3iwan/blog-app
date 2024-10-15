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

	DB, err := pgxpool.New(context.Background(),
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			postgresCfg.Host, postgresCfg.Port,
			postgresCfg.User, postgresCfg.Password,
			postgresCfg.DBName,
		),
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to the database: %v", err))
	}

	log.Println("Database connected successfully")
	return DB, nil
}
