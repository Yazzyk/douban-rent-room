package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/service/dataClean"
	"github.com/yazzyk/douban-rent-room/internal/service/notice"
	"github.com/yazzyk/douban-rent-room/internal/service/spider"
	"github.com/yazzyk/douban-rent-room/pkg/log"
)

func main() {
	config.Setup()
	log.SetupLogger()
	c := cron.New(cron.WithSeconds())

	for _, cron := range config.App.Cron {
		c.AddFunc(cron, func() {
			logrus.Info("Run")
			notice.Run(dataClean.Run(spider.Run()))
			logrus.Info("End")
		})
		c.Start()
	}

	fmt.Println("=== Start ===")
	select {}
}
