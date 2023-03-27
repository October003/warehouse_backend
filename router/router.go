package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"strconv"
	"warehouse/common/middleware"
)

// 需要角色鉴权的路由
var (
	//routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole = make([]func(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware), 0)
)

func Init() {
	r := gin.Default()
	//初始化 路由 中间件
	middleware.InitMiddleware(r)
	auth, err := middleware.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error,%s", err.Error())
	}
	// 注册 路由
	InitSysRouter(r, auth)
	addr := config.ApplicationConfig.Host + ":" + strconv.Itoa(int(config.ApplicationConfig.Port))
	r.Run(addr)
}
