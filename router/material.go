package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, SysMaterialRouter)
}

func SysMaterialRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	api := apis.Material{}
	r := v1.Group("/item").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/create", api.Create)
		r.GET("/get", api.Query)
		// ???runtime error: invalid memory address or nil pointer dereference
		r.DELETE("/delete", api.Delete)
		r.PUT("/update", api.Update)
	}
}
