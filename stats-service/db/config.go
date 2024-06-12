package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mdportnov/common/util"
	"log"
	"time"
)

var DB *pgxpool.Pool

func ConnectDatabase() {
	dbUser := util.GetEnv("DB_USER", "user")
	dbPassword := util.GetEnv("DB_PASSWORD", "password")
	dbName := util.GetEnv("DB_NAME", "nba")
	dbHost := util.GetEnv("DB_HOST", "localhost")
	dbPort := util.GetEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = DB.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	log.Println("Connected to the database")
}
