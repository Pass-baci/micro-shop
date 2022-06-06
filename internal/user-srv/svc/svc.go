package svc

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"micro-shop/internal/pkg/logger"
	"micro-shop/internal/user-srv/config"
	"micro-shop/internal/user-srv/store"
)

type Svc struct {
	Config *config.Config
	Logger logger.LoggerInterface
	Store  store.Querier
}

func NewSvc(configInfo *config.Config) *Svc {
	//_, err := db.NewMysqlDB(&db.MysqlInfo{
	//	Host:     configInfo.Mysql.Host,
	//	Port:     configInfo.Mysql.Port,
	//	Database: configInfo.Mysql.Database,
	//	Username: configInfo.Mysql.Username,
	//	Password: configInfo.Mysql.Password,
	//}, db.WithModel(&model.User{}))
	//if err != nil {
	//	log.Fatalf("连接数据库失败：%s", err.Error())
	//}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configInfo.Mysql.Username, configInfo.Mysql.Password, configInfo.Mysql.Host, configInfo.Mysql.Port, configInfo.Mysql.Database)

	db, err := sql.Open("mysql", dsn)
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
		Store: store.New(db),
	}
}
