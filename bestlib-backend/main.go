package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    _ "github.com/lib/pq"
    "github.com/gofiber/fiber/v2"
    "bestlib-backend/config"
    "bestlib-backend/routes"
)

func main() {
    config.InitElasticClient()

    connStr := "user=postgres dbname=bestlib sslmode=disable" 
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }
    defer db.Close()

    err = executeSQLScript(db, "setup.sql")
    if err != nil {
        log.Fatalf("Error executing SQL script: %v", err)
    }

    app := fiber.New()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}

func executeSQLScript(db *sql.DB, filename string) error {
    script, err := ioutil.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("error reading SQL file: %v", err)
    }

    _, err = db.Exec(string(script))
    if err != nil {
        return fmt.Errorf("error executing SQL script: %v", err)
    }

    return nil
}

