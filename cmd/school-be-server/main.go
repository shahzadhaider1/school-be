package main

import (
	"school-be/internal/api"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	handler := api.NewHandler()

	api.SetUpRoutes(app, handler)

	app.Listen(":3000")
}
