package config

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/config/source/file"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"warehouse/common/database"
)

func Init() {
	// 配置数据库 -- 初始化数据库
	config.Setup(
		file.NewSource(file.WithPath("config/settings.yml")),
		database.Setup,
	)

	application, errs := json.MarshalIndent(config.ApplicationConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		fmt.Println("application:", string(application))
	}

	jwt, errs := json.MarshalIndent(config.JwtConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		fmt.Println("jwt:", string(jwt))
	}
	fmt.Println("jwt:", string(jwt))
	database, errs := json.MarshalIndent(config.DatabasesConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		fmt.Println("database:", string(database))
	}

	loggerConfig, errs := json.MarshalIndent(config.LoggerConfig, "", "   ") //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
		fmt.Println("logger:", string(loggerConfig))
	}
}
