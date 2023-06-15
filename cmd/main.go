package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/api"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/db/bolt"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"github.com/yazzyk/douban-rent-room/internal/service/dataClean"
	"github.com/yazzyk/douban-rent-room/internal/service/notice"
	"github.com/yazzyk/douban-rent-room/internal/service/spider"
	"github.com/yazzyk/douban-rent-room/pkg/log"
)

func main() {
	config.Setup()
	log.SetupLogger()
	bolt.Setup()
	c := cron.New(cron.WithSeconds())

	for _, cron := range config.App.Cron {
		c.AddFunc(cron, func() {
			logrus.Info("====== Run ======")
			var result []models.HouseInfo
			for _, url := range config.App.Spider.WebSite {
				result = append(result, spider.Run(url)...)
			}
			notice.Run(dataClean.Run(result))
			logrus.Info("====== End ======")
		})
		c.Start()
	}

	api.Routers()
	//select {}
}
