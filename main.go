package main

import (
	log "github.com/go-admin-team/go-admin-core/logger"
	"warehouse/common/database"
	"warehouse/config"
	"warehouse/router"
)

func main() {
	// 配置信息初始化
	if err := config.Init(); err != nil {
		log.Fatal("config.Init() failed ,err:", err)
	}
	// 数据库信息初始化
	if err := database.Setup(); err != nil {
		log.Fatal("database.Setup() failed,err:", err)
	}
	// 路由初始化
	if err := router.Init(); err != nil {
		log.Fatal("router.Init() failed,err:", err)
	}
}
