package main

import (
	"log"
	"micro-shop/internal/pkg/conf"
	"micro-shop/internal/user-srv/config"
	"micro-shop/internal/user-srv/svc"
)

func main() {
	var configInfo = &config.Config{}
	if err := conf.ResolveConfig(configInfo, "./etc", "config"); err != nil {
		log.Fatalf("读取配置文件失败：%s", err.Error())
	}
	svc.NewSvc(configInfo)
}
