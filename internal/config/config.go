package config

import (
	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	Cron      []string
	Spider    SpiderConfig
	DataClean DataCleanConfig
	Logger    LoggerConfig
	Api       struct {
		Domain string
		Port   int
	}
	Notice struct {
		WxPusher struct {
			Enable   bool
			AppToken string
			Uids     []string
		}
		Email struct {
			Enable   bool
			From     string
			To       []string
			User     string
			Pwd      string
			SmtpAddr string
			Host     string
		}
	}
}

type SpiderConfig struct {
	WebSite   string
	TimeLimit int
	Cookie    string
}

type DataCleanConfig struct {
	//DistrictList []string
	BlackList []string
}

type LoggerConfig struct {
	Formatter string
	Level     string
	Path      string
}

var App AppConfig

func Setup() {
	if _, err := toml.DecodeFile("config.toml", &App); err != nil {
		panic("配置文件解析错误:" + err.Error())
		return
	}
}
