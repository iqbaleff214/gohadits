package main

import (
	"embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iqbaleff214/gohadits/domain/hadith"
)

//go:embed data/tafsirq/*.json
var data embed.FS

func main() {
	app := setup()
	log.Fatal(app.Listen(":8000"))
}

// app setup
func setup() *fiber.App {
	// repository instantiate
	hadithRepo := hadith.NewRepository(data)

	// service instantiate
	hadithService := hadith.NewService(hadithRepo)

	// handler instantiate
	hadithHandler := hadith.NewHandler(hadithService)

	app := fiber.New(config())
	app.Use(cors.New())

	// route
	api := app.Group("/api")

	// api v1
	api.Get("/v1", rootV1Handler)

	// hadith v1 domain
	api.Get("/v1/hadith", hadithHandler.GetAvailableBooks)
	api.Get("/v1/hadith/:book", hadithHandler.GetHadithByBook)
	api.Get("/v1/hadith/:book/:number", hadithHandler.GetHadithByNumber)

	return app
}

// route configuration
func config() fiber.Config {
	return fiber.Config{
		AppName:      "GoHadits",
		ErrorHandler: errorHandler,
	}
}
