package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, SysUserRouter)
}
func SysUserRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(auth.MiddlewareFunc()).Use(actions.PermissionAction())
	{
		r.GET("/get", api.GetPage)
		r.GET("/get/:id", api.Get)
		r.POST("/create", api.Insert)
		r.PUT("/update", api.Update)
		// ??? error: 34bfb972-059f-4ffb-8753-426f2d433ca9 无权删除该数据
		r.DELETE("/delete/:id", api.Delete)
	}
	r1 := v1.Group("/user").Use(auth.MiddlewareFunc()).Use(actions.PermissionAction())
	{
		r1.PUT("/pwd/reset", api.ResetPwd)
	}
}
