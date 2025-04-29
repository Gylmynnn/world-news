package router

import (
	"github.com/Gylmynnn/world-news/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRouter() error {
	app := fiber.New()

	app.Get("/scrape/:country", controller.ScrapeNews)

	return app.Listen(":3000")
}
