package controller

import (
	"github.com/Gylmynnn/world-news/service"
	"github.com/Gylmynnn/world-news/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Welcome(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(
		util.ResFormatter{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Message:    "welkam, hallo world",
			Data: fiber.Map{
				"endpoint":            "https://miniature-kellina-dice24434-c5ba5e44.koyeb.app/scrape/:country?limit=10&page=1",
				"countries available": [3]string{"indonesian", "japan", "chinese"},
			},
		},
	)

}

func ScrapeNews(c *fiber.Ctx) error {
	country := c.Params("country")
	limitParam := c.Query("limit", "10")
	pageParam := c.Query("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.ResFormatter{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
				Message:    fiber.ErrBadRequest.Message,
				Data:       err.Error(),
			},
		)
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.ResFormatter{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
				Message:    fiber.ErrBadRequest.Message,
				Data:       err.Error(),
			},
		)
	}

	newList, err := service.ScrapNewsByCountry(country, limit, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			util.ResFormatter{
				Success:    false,
				StatusCode: fiber.StatusInternalServerError,
				Message:    fiber.ErrInternalServerError.Message,
				Data:       err.Error(),
			},
		)
	}

	if len(newList) < 1 {
		return c.Status(fiber.StatusNotFound).JSON(
			util.ResFormatter{
				Success:    false,
				StatusCode: fiber.StatusNotFound,
				Message:    fiber.ErrNotFound.Message,
				Data:       newList,
			},
		)

	}

	return c.Status(fiber.StatusOK).JSON(
		util.ResFormatter{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Message:    "successfully",
			Data:       newList,
		},
	)
}
