package middleware

import (
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"net/http"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		//PayLoad 负载 是一个JSON对象 用来存放实际需要传递的数据
		// value, exists = c.Keys[key]
		// Keys is a key/value pair exclusively for the context of each request.键是专门用于每个请求上下文的键/值对。
		// Keys map[string]interface{}
		// const JwtPayloadKey = "JWT_PAYLOAD"
		// data ,_ = c.Keys["JWT_PAYLOAD"]
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		// 将data 装换为 map[string]interface{} 类型
		v := data.(jwtauth.MapClaims)
		// For client requests, Host optionally overrides the Host header to send. If empty,
		// the Request.Write method uses the value of URL.Host. Host may contain an international domain name.
		// 对于客户端请求，主机可以选择覆盖主机要发送的标头。如果为空，则 Request.Write 方法使用网址的值。主机。主机可能包含国际域名。

		// 执行器 Enforcer 是Casbin的主要结构。 它是用户就规则和模式开展业务的一个接口。
		// SyncedEnforer 基于 Enforcer() 并提供同步访问。 它是线程安全的。
		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		var res, casbinExclude bool
		var err error
		//检查权限
		if v["rolekey"] == "admin" {
			res = true
			// Next should be used only inside middleware. It executes the pending handlers in the chain inside the calling handler.
			c.Next()
			return
		}
		//遍历路由列表
		for _, i := range CasbinExclude {
			// bool function_name(string url, string pattern) 它返回一个布尔值表示 url 是否匹配 模式。
			// 匹配路由表中的  url 和 method
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		// Enforce决定“主体”是否能够用“动作”操作访问“对象”，输入参数通常是: (sub, obj, act)。
		res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, 500, err, "")
			return
		}
		if res {
			log.Infof("isTrue: %v role: %s method: %s path: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}

	}
}
