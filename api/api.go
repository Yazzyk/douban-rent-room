package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yazzyk/douban-rent-room/api/info"
	"github.com/yazzyk/douban-rent-room/internal/config"
)

func Routers() {
	app := fiber.New(
		fiber.Config{AppName: "豆瓣租房"},
	)
	app.Use(cors.New())

	app.Mount("/info", info.Router())

	app.Listen(fmt.Sprintf(":%d", config.App.Api.Port))
}
