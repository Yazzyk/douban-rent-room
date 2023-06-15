package info

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
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
	var result []models.HouseInfo
	for _, url := range config.App.Spider.WebSite {
		result = append(result, spider.Run(url)...)
	}
	notice.Run(dataClean.Run(result))
	logrus.Info("====== API End ======")
	return c.SendStatus(http.StatusOK)
}
