package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/middleware"
	"warehouse/common/middleware/handler"
)

func InitSysRouter(r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	g := r.Group("")
	// 不需要 登录认证 角色认证 的路由组
	NoCheckRoleRouterInit(g, auth)
	// 需要角色认证的路由组
	CheckRoleRouterInit(g, auth)
}

func NoCheckRoleRouterInit(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("/api")
	for _, f := range routerCheckRole {
		f(v1, auth)
	}
}

func CheckRoleRouterInit(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("/api/user")
	{
		v1.POST("/login", auth.LoginHandler)
		// Refresh time can be longer than token timeout
		v1.POST("/logout", handler.LogOut)
		v1.GET("/refresh_token", auth.RefreshHandler)
	}
	registerBaseRouter(v1, auth)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}
	api2 := apis.SysDept{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		//v1.GET("/menuTreeselect", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeselect/:roleId", api2.GetDeptTreeRoleSelect)
		//v1auth.POST("/logout", handler.LogOut)
	}
}
