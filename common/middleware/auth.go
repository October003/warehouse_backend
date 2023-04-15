package middleware

import (
	"time"
	"warehouse/common/middleware/handler"

	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"warehouse/config"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	if config.Cfg.Application.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.Cfg.Jwt.Timeout != 0 {
			timeout = time.Duration(config.Cfg.Jwt.Timeout) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte(config.Cfg.Jwt.Secret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc, //负载
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
