package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"strconv"
	"warehouse/common/middleware"
	"warehouse/config"
)

// 需要角色鉴权的路由
var (
	//routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole = make([]func(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware), 0)
)

func Init() (err error) {
	r := gin.Default()
	//gin.SetMode(gin.TestMode)
	//gin.SetMode(gin.ReleaseMode)
	//初始化 路由 中间件
	middleware.InitMiddleware(r)
	auth, err := middleware.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error,%s", err.Error())
		return err
	}
	// 注册 路由
	InitSysRouter(r, auth)
	Addr := config.Cfg.Application.Host + ":" + strconv.Itoa(int(config.Cfg.Application.Port))
	r.Run(Addr)
	return nil
}
