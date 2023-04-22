package database

import (
	"bufio"
	"bytes"
	_ "database/sql"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"warehouse/config"

	mycasbin "github.com/go-admin-team/go-admin-core/sdk/pkg/casbin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"sync"
	models2 "warehouse/common/models"
	"warehouse/models"
)

var once sync.Once
var DB *gorm.DB

func Setup() (err error) {
	dsn := config.Cfg.Database.Source
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("connect to mysql failed")
		return err
	}
	host := "*"
	e := mycasbin.Setup(DB, "")
	sdk.Runtime.SetDb(host, DB)
	sdk.Runtime.SetCasbin(host, e)
	err = DB.Migrator().AutoMigrate(
		&models.ItemSql{},
		&models.InboundSql{},
		&models.OutboundSql{},
		&models.InBoundPersons{},
		&models.OutBoundPersons{},
		&models.Units{},
		&models.StrongLocation{},
		&models.ItemIDs{},
		&models2.SysDept{},
		&models2.SysMenu{},
		&models2.SysRoleDept{},
		&models2.SysUser{},
		&models2.SysRole{},
		&models2.SysApi{},
	)
	if err != nil {
		log.Fatal("db.Migrator().AutoMigrate() failed")
		return err
	}
	once.Do(func() {
		var err error
		// 读取 SQL 文件
		sqlBytes, err := ioutil.ReadFile("config/db.sql")
		if err != nil {
			log.Fatal(err)
		}
		// 拆分 SQL 语句并执行
		scanner := bufio.NewScanner(bytes.NewReader(sqlBytes))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 || line[0] == '#' {
				continue // 忽略空行和注释行
			}
			err = DB.Exec(line).Error
			if err != nil {
				log.Info("Error executing SQL statement: %s\n%s\n", err, line)
				continue
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	})
	return nil
}
