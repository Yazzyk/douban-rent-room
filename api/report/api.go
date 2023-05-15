package report

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yazzyk/douban-rent-room/internal/service/report"
	"net/http"
)

func Router() *fiber.App {
	api := fiber.New()
	{
		api.Get("", reportUser)
	}
	return api
}

func reportUser(c *fiber.Ctx) error {
	id := c.Query("id")
	name := c.Query("name")
	if id == "" || name == "" {
		return c.Status(http.StatusBadRequest).SendString("请求错误")
	}

	report.User(id, name)

	return nil
}
