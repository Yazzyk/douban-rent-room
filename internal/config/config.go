package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type AppConfig struct {
	Cron      []string
	Sort      string
	Spider    SpiderConfig
	DataClean DataCleanConfig
	Logger    LoggerConfig
	DB        struct {
		FileName string
	}
	Api struct {
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
	CookieCloud CookieCloudConfig
}

type CookieCloudConfig struct {
	ServerHost string
	UUID       string
	Password   string
}

type SpiderConfig struct {
	WebSite   []string
	TimeLimit int
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
	confFile := "config.toml"
	if len(os.Args) >= 2 {
		confFile = fmt.Sprintf("config_%s.toml", os.Args[1])
	}
	if _, err := toml.DecodeFile(confFile, &App); err != nil {
		panic("配置文件解析错误:" + err.Error())
		return
	}
}
