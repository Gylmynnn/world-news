package router

import (
	"os"

	"github.com/Gylmynnn/world-news/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRouter() error {
	app := fiber.New()

	app.Get("/", controller.Welcome)

	app.Get("/scrape/:country", controller.ScrapeNews)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return app.Listen(":" + port)
}
