package main

import (
	"school-be/internal/api"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	api.SetUpRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
