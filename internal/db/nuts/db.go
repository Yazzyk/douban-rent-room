package nuts

import (
	"github.com/nutsdb/nutsdb"
	"github.com/sirupsen/logrus"
)

var DB *nutsdb.DB

func Setup() {
	var err error
	DB, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("nutsDB"),
	)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func Put(bucket, key, value string) (err error) {
	if err = DB.Update(func(tx *nutsdb.Tx) error {
		if err = tx.Put(bucket, []byte(key), []byte(value), 0); err != nil {
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
	if err := DB.View(
		func(tx *nutsdb.Tx) error {
			e, err := tx.Get(bucket, []byte(key))
			if err != nil {
				return err
			}
			value = string(e.Value)
			return nil
		}); err != nil {
		logrus.Error(err)
	}
	return
}

func View(bucket string) (result map[string]string) {
	result = make(map[string]string)
	if err := DB.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll(bucket)
			if err != nil {
				return err
			}

			for _, entry := range entries {
				result[string(entry.Key)] = string(entry.Value)
			}
			return nil
		}); err != nil {
		logrus.Error(err)
	}
	return
}
