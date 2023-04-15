package dataClean

import (
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"strings"
)

func Run(dataList []models.HouseInfo) (result []models.HouseInfo) {
	logrus.Info("数据清洗")
LIST:
	for _, info := range dataList {
		// 清洗黑名单
		for _, s := range config.App.DataClean.BlackList {
			if strings.Contains(info.Title, s) {
				continue LIST
			}
		}
		result = append(result, info)
	}
	logrus.Info("数据清洗完成")
	return
}
