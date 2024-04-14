package api

import (
	"github.com/gofiber/fiber/v3"

	"school-be/internal/service"
)

type Handler struct {
	CustomerService *service.CustomerService
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hello, From Pluckers 👋!"})
}

func (h *Handler) GetAllNewsAndOffers(c *fiber.Ctx) error {
	return nil
}
