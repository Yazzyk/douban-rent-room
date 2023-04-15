package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"io"
	"os"
	"runtime"
	"time"
)

func SetupLogger() {

	switch config.App.Logger.Formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&nested.Formatter{
			CallerFirst:     true,
			HideKeys:        true,
			FieldsOrder:     []string{"component", "category"},
			TimestampFormat: time.DateTime,
			CustomCallerFormatter: func(f *runtime.Frame) string {
				return fmt.Sprintf(" [ %s:%d ] ", f.File, f.Line)
			},
		})
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, newFileWriter(config.App.Logger.Path)))
	l, err := logrus.ParseLevel(config.App.Logger.Level)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.SetLevel(l)
	logrus.SetReportCaller(true)
}

func newFileWriter(filename string) *rotatelogs.RotateLogs {
	path := filename + ".%Y%m%d.log"
	writer, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return writer
}
