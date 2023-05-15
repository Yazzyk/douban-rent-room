package report

import (
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/db/bolt"
)

const (
	bucket = "report"
)

func User(id, name string) {
	bolt.Put(bucket, id, name)
	logrus.Info("屏蔽用户:", name)
}
