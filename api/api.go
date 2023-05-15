package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yazzyk/douban-rent-room/api/info"
	"github.com/yazzyk/douban-rent-room/api/report"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"time"
)

func Routers() {
	app := fiber.New(
		fiber.Config{AppName: "豆瓣租房"},
	)
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: fmt.Sprintf("%s ${status} ${method} ${path} \n", time.Now().Format("2006-01-02 15:04:05")),
	}))

	app.Mount("/info", info.Router())
	app.Mount("/report", report.Router())

	app.Listen(fmt.Sprintf(":%d", config.App.Api.Port))
}
