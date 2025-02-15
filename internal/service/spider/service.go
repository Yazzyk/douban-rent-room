package spider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/cookieCloudSDK"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run(website string) (result []models.HouseInfo) {
	logrus.Info("开始爬取数据")
	start := 0
	endTime := time.Now().Add(-time.Duration(config.App.Spider.TimeLimit) * 24 * time.Hour)
	cookie, err := GetCookie()
	if err != nil {
		logrus.Error(err)
		return
	}
	for {
		pageResp, err := resty.New().R().SetHeader("Cookie", cookie).Get(fmt.Sprintf("%s?start=%d&type=new", website, start))
		if err != nil {
			logrus.Error(err)
			return
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(pageResp.Body()))
		if err != nil {
			logrus.Error(err)
			return
		}

		webTitle := doc.Find("title").First().Text()

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

			result = append(result, models.HouseInfo{
				Title:        title,
				Link:         link,
				Author:       userName,
				AuthorLink:   userLinkStr,
				AuthorID:     userID,
				Date:         &date,
				DataFrom:     webTitle,
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
			logrus.Infof("[%s]共有%d条", webTitle, len(result))
			return
		}
		start += 30
	}
}

func GetCookie() (cookieStr string, err error) {
	cc, err := cookieCloudSDK.NewCookieCloudSDK(config.App.CookieCloud.ServerHost, config.App.CookieCloud.UUID, config.App.CookieCloud.Password)
	if err != nil {
		logrus.Error(err)
		return
	}
	cookieData, err := cc.GetCookie()
	if err != nil {
		logrus.Error(err)
		return
	}
	//logrus.Info("从CookieCloud获取到Cookie: ", cookieData)
	data, exist := cookieData.CookieData["douban.com"]
	if !exist {
		logrus.Warn("未获取到[douban.com]的cookie,尝试www.douban.com")
		data, exist = cookieData.CookieData["www.douban.com"]
		if !exist {
			logrus.Error("未获取到Cookie")
		}
		return
	}
	for _, datum := range data {
		cookieStr += fmt.Sprintf("%s=%s;", datum.Name, datum.Value)
	}
	logrus.Info("从CookieCloud获取到Cookie成功: ", cookieStr)
	return
}
