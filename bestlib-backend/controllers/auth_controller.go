package controllers

import (
    "github.com/gofiber/fiber/v2"
    "bestlib-backend/models"
    "bestlib-backend/services"
    "net/http"
)

type AuthController struct {
    AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{AuthService: authService}
}

func (ctrl *AuthController) Register(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    if err := ctrl.AuthService.RegisterUser(&user); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    token, err := ctrl.AuthService.LoginUser(user.IIN, user.Password)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"token": token})
}
