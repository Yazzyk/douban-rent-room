package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
	"github.com/yazzyk/douban-rent-room/internal/config"
	"os"
	"path/filepath"
)

var db *bolt.DB

const (
	REPORT_BUCKET = "report"
)

func Setup() {
	_, err := os.Stat(filepath.FromSlash(config.App.DB.FileName))
	if err != nil {
		if err = os.MkdirAll(filepath.FromSlash(filepath.Dir(config.App.DB.FileName)), os.ModePerm); err != nil {
			logrus.Error(err)
		}

	}

	d, err := bolt.Open(config.App.DB.FileName, os.ModePerm, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	db = d

	buckets()
}

func buckets() {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(REPORT_BUCKET))
		if err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	})
}

func Put(bucket, key, value string) (err error) {
	if err = db.Update(func(tx *bolt.Tx) error {
		if err = tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(value)); err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	}); err != nil {
		logrus.Error(err)
		return
	}
	return
}

func Get(bucket, key string) (value string) {
	if err := db.View(
		func(tx *bolt.Tx) error {
			value = string(tx.Bucket([]byte(bucket)).Get([]byte(key)))
			return nil
		}); err != nil {
		logrus.Error(err)
	}
	return
}

func View(bucket string) (result map[string]string) {
	result = make(map[string]string)
	if err := db.View(
		func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte(bucket)).ForEach(func(k, v []byte) error {
				result[string(k)] = string(v)
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
		logrus.Error(err)
	}
	return
}
