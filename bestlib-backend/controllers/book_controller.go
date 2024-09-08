package controllers

import (
    "github.com/gofiber/fiber/v2"
    "bestlib-backend/services"
    "net/http"
)

type BookController struct {
    BookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
    return &BookController{BookService: bookService}
}

func (ctrl *BookController) SearchBooks(c *fiber.Ctx) error {
    searchTerm := c.Query("search_term")
    books, err := ctrl.BookService.SearchBooks(searchTerm)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(books)
}
