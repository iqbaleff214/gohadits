package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := setup()
	log.Fatal(app.Listen(":8000"))
}

func setup() *fiber.App {
	app := fiber.New(config())
	app.Use(cors.New())

	
	api := app.Group("/api/")

	// api version 1
	apiV1 := api.Group("/v1")

	apiV1.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]any{
			"owner":   "M. Iqbal Effendi <iqbaleff214@gmail.com>",
			"license": "MIT",
			"source":  "https://github.com/iqbaleff214/gohadits",
			"version": "1",
			"endpoints": []map[string]string{
				{
					"path":        "/api/v1/hadith",
					"method":      "GET",
					"description": "Mengembalikan daftar kitab hadits yang tersedia.",
				},
				{
					"path":        "/api/v1/hadith/{book}",
					"method":      "GET",
					"description": "Mengembalikan daftar hadits yang terdapat pada kitab yang diminta.",
				},
				{
					"path":        "/api/v1/hadith/{book}/{number}",
					"method":      "GET",
					"description": "Mengembalikan hadith spesifik.",
				},
			},
		})
	})

	return app
}

func config() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
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
		},
	}
}
