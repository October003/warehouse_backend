package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
//func Options(c *gin.Context) {
//	if c.Request.Method != "OPTIONS" {
//		c.Next()
//	} else {
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
//		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
//		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
//		c.Header("Content-Type", "application/json")
//		c.AbortWithStatus(200)
//	}
//}

func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin") //请求头部
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
	//允许类型校验
	if method == "OPTIONS" {
		c.JSON(http.StatusOK, "ok!")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	c.Next()
}
