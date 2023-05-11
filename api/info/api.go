package info

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/service/dataClean"
	"github.com/yazzyk/douban-rent-room/internal/service/notice"
	"github.com/yazzyk/douban-rent-room/internal/service/spider"
	"net/http"
)

func Router() *fiber.App {
	api := fiber.New()

	{
		api.Get("", getInfo)
	}

	return api
}

func getInfo(c *fiber.Ctx) error {
	logrus.Info("====== API Run ======")
	notice.Run(dataClean.Run(spider.Run()))
	logrus.Info("====== API End ======")
	return c.SendStatus(http.StatusOK)
}
