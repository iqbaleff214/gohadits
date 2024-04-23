package hadith

import "github.com/gofiber/fiber/v2"

// Handler contains method for fiber route handlers
type Handler interface {
	GetAvailableBooks(c *fiber.Ctx) error
}

// handler as a class
type handler struct {
	service Service
}

// NewHandler is a function to instantiate new handler object
func NewHandler(service Service) Handler {
	return handler{service}
}

// GET /api/v1/hadith
func (h handler) GetAvailableBooks(c *fiber.Ctx) error {
	books := h.service.GetAllAvailableBooks()

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"code":    fiber.StatusOK,
		"message": "Hadith books successfully retrieved.",
		"status":  "success",
		"data":    books,
	})
}

