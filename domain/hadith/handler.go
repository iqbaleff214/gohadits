package hadith

import "github.com/gofiber/fiber/v2"

const (
	DEFAULT_LIMIT_PER_PAGE int = 50
	MAX_LIMIT_PER_PAGE     int = 150
)

// Handler contains method for fiber route handlers
type Handler interface {
	GetAvailableBooks(c *fiber.Ctx) error
	GetHadithByBook(c *fiber.Ctx) error
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

// GET /api/v1/hadith/:book
func (h handler) GetHadithByBook(c *fiber.Ctx) error {
	bookName := c.Params("book")

	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", DEFAULT_LIMIT_PER_PAGE)

	book, hadiths, err := h.service.GetHadithByBook(bookName, offset, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"code":    fiber.StatusOK,
		"message": "Hadith collection of " + book.Name + " successfully retrieved.",
		"status":  "success",
		"data": map[string]any{
			"name":   book.Name,
			"total":  book.Size,
			"items":  hadiths,
			"offset": offset,
			"limit":  limit,
		},
	})
}
