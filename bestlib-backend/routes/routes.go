package routes

import (
    "github.com/gofiber/fiber/v2"
    "bestlib-backend/controllers"
    "bestlib-backend/repository"
    "bestlib-backend/services"
    "github.com/jackc/pgx/v5/pgxpool"
    "context"
    "log"
)

func SetupRoutes(app *fiber.App) {
    dsn := "postgresql://postgres:password@localhost:5432/bestlib"
    dbpool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    userRepo := repository.NewUserRepository(dbpool) 
    bookRepo := repository.NewBookRepository(dbpool)
    authService := services.NewAuthService(userRepo)
    bookService := services.NewBookService(bookRepo)

    authController := controllers.NewAuthController(authService)
    bookController := controllers.NewBookController(bookService)

    app.Post("/register", authController.Register)
    app.Post("/login", authController.Login)
    app.Get("/search", bookController.SearchBooks)
}

