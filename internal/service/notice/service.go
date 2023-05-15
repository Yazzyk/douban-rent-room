package notice

import (
	"bytes"
	_ "embed"
	"github.com/PaleBlueYk/wxpusher-sdk-go"
	"github.com/PaleBlueYk/wxpusher-sdk-go/msg"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"github.com/yazzyk/douban-rent-room/internal/models"
	"html/template"
	"net/smtp"
)

//go:embed notice.html
var noticeTemplate string

func Run(data []models.HouseInfo) {
	if config.App.Notice.WxPusher.Enable {
		logrus.Info("微信推送")
		wxPusherSender(data)
	}
	if config.App.Notice.Email.Enable {
		logrus.Info("邮件发送")
		emailSender(data)
	}
}

func wxPusherSender(data []models.HouseInfo) {
	tmpl, err := template.New("notice").Parse(noticeTemplate)
	if err != nil {
		logrus.Error(err)
		return
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]interface{}{
		"Data":  data,
		"Black": config.App.DataClean.BlackList,
		"Day":   config.App.Spider.TimeLimit,
	})
	m := msg.NewMessage(config.App.Notice.WxPusher.AppToken).SetContent(buf.String()).AddUId("", config.App.Notice.WxPusher.Uids...)
	message, err := wxpusher.SendMessage(m)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(message)
}

func emailSender(data []models.HouseInfo) {
	mail := email.NewEmail()
	mail.From = config.App.Notice.Email.From
	mail.To = config.App.Notice.Email.To
	mail.Subject = "豆瓣租房信息"
	tmpl, err := template.New("notice").Parse(noticeTemplate)
	if err != nil {
		logrus.Error(err)
		return
	}
	var buf bytes.Buffer
	t := make(map[string]interface{})
	t = map[string]interface{}{
		"Data":   data,
		"Black":  config.App.DataClean.BlackList,
		"Day":    config.App.Spider.TimeLimit,
		"Domain": config.App.Api.Domain,
	}
	err = tmpl.Execute(&buf, t)
	if err != nil {
		logrus.Error(err)
		return
	}
	mail.HTML = buf.Bytes()
	err = mail.Send(config.App.Notice.Email.SmtpAddr, smtp.PlainAuth("", config.App.Notice.Email.User, config.App.Notice.Email.Pwd, config.App.Notice.Email.Host))
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("邮件发送成功")
}
