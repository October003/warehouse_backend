package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

// 需认证的路由代码
func registerSysDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysDept{}

	r := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/get", api.GetPage)
		r.GET("/get/:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update/:id", api.Update)
		r.DELETE("/delete", api.Delete)
	}
	r1 := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r1.GET("/deptTree", api.Get2Tree)
	}
}
