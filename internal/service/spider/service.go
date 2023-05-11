package spider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"os"
	"time"
)

func Run() (result []models.HouseInfo) {
	logrus.Info("开始爬取数据")
	start := 0
	endTime := time.Now().Add(-time.Duration(config.App.Spider.TimeLimit) * 24 * time.Hour)
	for {
		pageResp, err := resty.New().R().SetHeader("Cookie", config.App.Spider.Cookie).Get(fmt.Sprintf("%s?start=%d&type=new", config.App.Spider.WebSite, start))
		if err != nil {
			logrus.Error(err)
			return
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(pageResp.Body()))
		if err != nil {
			logrus.Error(err)
			return
		}

		doc.Find("table.olt tr").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				return
			}
			link, _ := selection.Find("td.title").First().Find("a").Attr("href")
			title, _ := selection.Find("td.title").First().Find("a").Attr("title")
			date, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%d-%s", time.Now().Year(), selection.Find("td").Last().Text()))
			if err != nil {
				logrus.Error(err)
				return
			}
			result = append(result, models.HouseInfo{
				Title: title,
				Link:  link,
				Date:  &date,
			})
		})

		if len(result) == 0 {
			logrus.Error("未获取到数据")
			os.WriteFile("logs/index.html", []byte(doc.Text()), os.ModePerm)
			return
		}

		if len(result) != 0 && result[len(result)-1].Date.Unix() < endTime.Unix() {
			logrus.Infof("共有%d条", len(result))
			return
		}
		start += 30
	}
}
