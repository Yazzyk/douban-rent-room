package dataSort

import (
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"sort"
)

func Sort(data []models.HouseInfo) (result []models.HouseInfo) {
	logrus.Info("=== 数据排序 ===")
	switch config.App.Sort {
	case "time":
		logrus.Info("根据时间排序")
		return timeSort(data)
	default:
		logrus.Info("默认排序")
		return data
	}
}

func timeSort(data []models.HouseInfo) (result []models.HouseInfo) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Unix() > data[j].Date.Unix()
	})
	return data
}
