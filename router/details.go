package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/apis"
	"warehouse/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, SysDetailsRouter)
}
func SysDetailsRouter(v1 *gin.RouterGroup, auth *jwt.GinJWTMiddleware) {
	api := apis.Details{}
	r := v1.Group("/inbound").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/create", api.CreateInbound)
		r.DELETE("/delete", api.DeleteInbound)
		r.GET("/get", api.QueryInbound)
	}
	r1 := v1.Group("outbound").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.POST("/create", api.CreateOutbound)
		r1.DELETE("/delete", api.DeleteOutbound)
		r1.GET("/get", api.QueryOutbound)
	}
	r2 := v1.Group("query_by_timestamp").Use(auth.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r2.POST("", api.QueryByTimestamp)
	}
}
