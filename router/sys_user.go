package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/actions"
	"warehouse/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, SysUserRouter)
}

func SysUserRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(auth.MiddlewareFunc()).Use(actions.PermissionAction()).Use(middleware.AuthCheckRole())
	{
		r.GET("/get", api.GetPage)
		r.GET("/get/:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update", api.Update)
		r.DELETE("/delete/:id", api.Delete)
	}
	r1 := v1.Group("/user").Use(auth.MiddlewareFunc()).Use(actions.PermissionAction())
	{
		r1.PUT("/pwd/reset", api.ResetPwd)
	}
}
