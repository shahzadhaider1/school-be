package api

import (
	"strconv"

	"github.com/gofiber/fiber/v3"

	"school-be/internal/service"
	"school-be/pkg/models"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(classService *service.ClassService) *ClassHandler {
	return &ClassHandler{
		classService: classService,
	}
}

// CreateClassHandler handles the creation of a new class
func (h *ClassHandler) CreateClassHandler(c fiber.Ctx) error {
	class := new(models.Class)
	if err := c.Bind().Body(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing request body",
		})
	}

	if err := h.classService.CreateClass(class); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating class"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Class created successfully", "class": class})
}

// GetClassByIDHandler handles the retrieval of a class by its ID
func (h *ClassHandler) GetClassByIDHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid class ID"})
	}

	class, err := h.classService.GetClassByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Class not found"})
	}

	return c.JSON(fiber.Map{"class": class})
}

// ListClassesHandler handles the listing of all classes
func (h *ClassHandler) ListClassesHandler(c fiber.Ctx) error {
	filter := make(map[string]interface{})
	// You can parse query parameters or request body to apply filtering if needed

	classes, err := h.classService.ListClasses(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error listing classes"})
	}

	return c.JSON(fiber.Map{"classes": classes})
}

// UpdateClassHandler handles the update of an existing class
func (h *ClassHandler) UpdateClassHandler(c fiber.Ctx) error {
	class := new(models.Class)
	if err := c.Bind().Body(class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}

	if err := h.classService.UpdateClass(class); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating class"})
	}

	return c.JSON(fiber.Map{"message": "Class updated successfully", "class": class})
}

// DeleteClassHandler handles the deletion of a class by its ID
func (h *ClassHandler) DeleteClassHandler(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid class ID"})
	}

	if err := h.classService.DeleteClass(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting class"})
	}

	return c.JSON(fiber.Map{"message": "Class deleted successfully"})
}
