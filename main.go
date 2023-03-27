package main

import (
	"warehouse/config"
	"warehouse/router"
)

func main() {
	// 配置信息初始化
	config.Init()
	// 路由初始化
	router.Init()
}
