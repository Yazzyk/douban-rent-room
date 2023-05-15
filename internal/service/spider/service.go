package spider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/db/bolt"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run() (result []models.HouseInfo) {
	logrus.Info("开始爬取数据")
	start := 0
	endTime := time.Now().Add(-time.Duration(config.App.Spider.TimeLimit) * 24 * time.Hour)
	backUser := bolt.View("report")
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
			// 获取帖子信息
			link, _ := selection.Find("td.title").First().Find("a").Attr("href")
			title, _ := selection.Find("td.title").First().Find("a").Attr("title")
			date, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%d-%s", time.Now().Year(), selection.Find("td").Last().Text()))
			count, _ := strconv.ParseInt(selection.Find("td.r-count").First().Text(), 10, 64)
			user := selection.Find("td").Eq(1).Find("a")
			userLinkStr, _ := user.Attr("href")
			userName := user.Text()
			userLinkList := strings.Split(userLinkStr, "/")
			userID := userLinkList[len(userLinkList)-2]

			// 黑名单用户检查
			name, exist := backUser[userID]
			if exist {
				logrus.Warnf("用户[ %s ]已被屏蔽，跳过其发布信息", name)
				return
			}

			result = append(result, models.HouseInfo{
				Title:        title,
				Link:         link,
				Author:       userName,
				AuthorLink:   userLinkStr,
				AuthorID:     userID,
				Date:         &date,
				DateStr:      date.Format("2006-01-02 15:04"),
				CommentCount: int(count),
			})
		})

		if len(result) == 0 {
			logrus.Error("未获取到数据")
			os.WriteFile("logs/index.html", pageResp.Body(), os.ModePerm)
			return
		}

		if len(result) != 0 && result[len(result)-1].Date.Unix() < endTime.Unix() {
			logrus.Infof("共有%d条", len(result))
			return
		}
		start += 30
	}
}
