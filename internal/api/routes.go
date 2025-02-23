package api

import (
	"school-be/internal/service"

	"github.com/gofiber/fiber/v3"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")

	classService := service.NewClassService()
	classHandler := NewClassHandler(classService)

	api.Post("/classes", classHandler.CreateClassHandler)
	api.Get("/classes/:id", classHandler.GetClassByIDHandler)
	api.Get("/classes", classHandler.ListClassesHandler)
	api.Put("/classes/:id", classHandler.UpdateClassHandler)
	api.Delete("/classes/:id", classHandler.DeleteClassHandler)
}
