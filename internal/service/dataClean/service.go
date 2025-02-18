package dataClean

import (
	"context"
	"encoding/json"
	"github.com/coze-dev/coze-go"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/db/bolt"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"github.com/yazzyk/douban-rent-room/internal/service/spider"
	"strings"
)

func Run(dataList []models.HouseInfo) (result []models.HouseInfo) {
	logrus.Info("数据清洗，目前共有[ %d ]条数据", len(dataList))
	backUser := bolt.View("report")
LIST:
	for _, info := range dataList {
		// 黑名单用户检查
		name, exist := backUser[info.AuthorID]
		if exist {
			logrus.Warnf("用户[ %s ]已被屏蔽，跳过其发布信息", name)
			continue
		}
		// 清洗黑名单
		for _, s := range config.App.DataClean.BlackList {
			if strings.Contains(info.Title, s) {
				continue LIST
			}
		}
		result = append(result, info)
	}
	logrus.Info("黑民单数据清洗完成，目前共有[ %d ]条数据", len(result))
	var err error
	result, err = cozeAIClear(result)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("数据清洗完成")
	return
}

// ai 数据清洗
func cozeAIClear(dataList []models.HouseInfo) (result []models.HouseInfo, err error) {
	logrus.Info("AI数据清洗, 共有[ %d ]条数据", len(dataList))
	jwtOauthClient, err := coze.NewJWTOAuthClient(coze.NewJWTOAuthClientParam{
		ClientID:      config.App.DataClean.CozeClientID,
		PublicKey:     config.App.DataClean.CozePublicKey,
		PrivateKeyPEM: config.App.DataClean.CozePrivateKey,
	}, coze.WithAuthBaseURL("https://api.coze.cn"))
	if err != nil {
		logrus.Error(err)
		return
	}
	//tokenResp, err := jwtOauthClient.GetAccessToken(context.Background(), nil)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	//logrus.Info(tokenResp)
	cozeCli := coze.NewCozeAPI(coze.NewJWTAuth(jwtOauthClient, nil), coze.WithBaseURL("https://api.coze.cn"))
	cookie, err := spider.GetCookie()
	if err != nil {
		logrus.Error(err)
		cookie = config.App.Spider.Cookie
		if cookie == "" {
			logrus.Error("未获取到Cookie")
			return
		}
		err = nil
	}
	for _, info := range dataList {
		pageResp, pageErr := resty.New().R().SetHeader("Cookie", cookie).Get(info.Link)
		if pageErr != nil {
			logrus.Error(pageErr)
			continue
		}
		//logrus.Info(pageResp.String())
		resp, aierr := cozeCli.Workflows.Runs.Create(context.Background(), &coze.RunWorkflowsReq{
			WorkflowID: config.App.DataClean.WorkflowID,
			Parameters: map[string]interface{}{
				"spider_data":   pageResp.String(),
				"required_data": config.App.DataClean.Requirement,
			},
		})
		if aierr != nil {
			logrus.Error(aierr)
			continue
		}
		var aiResp aiResponse
		if err = json.Unmarshal([]byte(resp.Data), &aiResp); err != nil {
			logrus.Error(err)
			continue
		}
		logrus.Debug(resp.Data)
		if aiResp.Output {
			info.AIReason = aiResp.Reason
			info.CommutingTime = aiResp.Time
			result = append(result, info)
		}
	}
	logrus.Info("AI数据清洗完成, 共有[ %d ]条数据", len(result))
	return
}

type aiResponse struct {
	Output bool   `json:"output"`
	Reason string `json:"reason"`
	Time   string `json:"time"`
}
