package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, AdminRouter)
}

func AdminRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	api := apis.Admin{}
	r := v1.Group("config").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/get", api.Get)          //获取系统设置
		r.POST("/add", api.Add)         //添加系统设置
		r.DELETE("/delete", api.Delete) //删除系统设置
	}
}
