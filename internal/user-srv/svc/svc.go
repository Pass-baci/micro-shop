package svc

import (
	"log"
	"micro-shop/internal/model"
	"micro-shop/internal/pkg/conf"
	"micro-shop/internal/pkg/db"
	"micro-shop/internal/user-srv/config"
)

type Svc struct {
	Config *config.Config
}

func NewSvc() *Svc {
	var svcCtx = &Svc{Config: &config.Config{}}
	if err := conf.ResolveConfig(svcCtx.Config, "./etc", "config"); err != nil {
		log.Fatalf("读取配置文件失败：%s", err.Error())
	}
	configInfo := svcCtx.Config
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
	return svcCtx
}
