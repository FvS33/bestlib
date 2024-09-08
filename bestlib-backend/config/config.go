package config

import (
    "context"
    "fmt"
    "log"
    "github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func SetupDatabaseConnection() {
    dsn := "postgresql://postgres:password@localhost:5432/bestlib"
    var err error
    DBPool, err = pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    fmt.Println("Successfully connected to the database")
}

func GetDBPool() *pgxpool.Pool {
    return DBPool
}

