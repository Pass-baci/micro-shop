package svc

import (
	"log"
	"micro-shop/internal/model"
	"micro-shop/internal/pkg/db"
	"micro-shop/internal/pkg/logger"
	"micro-shop/internal/user-srv/config"
)

type Svc struct {
	Config *config.Config
	Logger logger.LoggerInterface
}

func NewSvc(configInfo *config.Config) *Svc {
	_, err := db.NewMysqlDB(&db.MysqlInfo{
		Host:     configInfo.Mysql.Host,
		Port:     configInfo.Mysql.Port,
		Database: configInfo.Mysql.Database,
		Username: configInfo.Mysql.Username,
		Password: configInfo.Mysql.Password,
	}, db.WithModel(&model.User{}))
	if err != nil {
		log.Fatalf("连接数据库失败：%s", err.Error())
	}
	return &Svc{
		Config: configInfo,
		Logger: logger.NewLogger(logger.LoggerInfo{
			Level:      configInfo.Mode,
			EncodeTime: configInfo.Logger.EncodeTime,
			LoggerFile: struct {
				LogPath    string
				ErrPath    string
				MaxSize    int
				MaxBackups int
				MaxAge     int
				Compress   bool
			}{
				LogPath:    configInfo.Logger.LoggerFile.LogPath,
				ErrPath:    configInfo.Logger.LoggerFile.ErrPath,
				MaxSize:    configInfo.Logger.LoggerFile.MaxSize,
				MaxBackups: configInfo.Logger.LoggerFile.MaxBackups,
				MaxAge:     configInfo.Logger.LoggerFile.MaxAge,
				Compress:   configInfo.Logger.LoggerFile.Compress,
			},
		}),
	}
}
