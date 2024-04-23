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
	api := app.Group("/api/")

	// api version 1
	apiV1 := api.Group("/v1")
	apiV1.Get("/", rootV1Handler)

	// hadith v1 domain
	hadithRoute := apiV1.Group("/hadith")
	hadithRoute.Get("/", hadithHandler.GetAvailableBooks)

	return app
}

// route configuration
func config() fiber.Config {
	return fiber.Config{
		AppName:      "GoHadits",
		ErrorHandler: errorHandler,
	}
}
