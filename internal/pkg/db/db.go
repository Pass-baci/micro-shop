package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInfo struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type option func(db *gorm.DB) error

func NewMysqlDB(conf *MysqlInfo, opts ...option) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt(db); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func WithModel(model ...interface{}) option {
	return func(db *gorm.DB) error {
		return db.AutoMigrate(model)
	}
}
