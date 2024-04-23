package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// GET /api/v1/
func rootV1Handler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"owner":   "M. Iqbal Effendi <iqbaleff214@gmail.com>",
		"license": "MIT",
		"source":  "https://github.com/iqbaleff214/gohadits",
		"version": "1",
		"endpoints": []map[string]string{
			{
				"path":        "/api/v1/hadith",
				"method":      "GET",
				"description": "Returning a list of available hadith books.",
			},
			{
				"path":        "/api/v1/hadith/{book}?offset=0&limit=50",
				"method":      "GET",
				"description": "Returning a list of hadiths found in the requested book.",
			},
			{
				"path":        "/api/v1/hadith/{book}/{number}",
				"method":      "GET",
				"description": "Returning a specific hadith according to the given hadith number.",
			},
		},
	})
}

// Error handler response
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = c.Status(code).JSON(map[string]any{
		"code":    code,
		"message": e.Message,
		"status":  "error",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal Server Error",
			"status":  "error",
		})
	}

	return nil
}
