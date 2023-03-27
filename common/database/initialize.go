package database

import (
	"bufio"
	"bytes"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	toolsConfig "github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	mycasbin "github.com/go-admin-team/go-admin-core/sdk/pkg/casbin"
	toolsDB "github.com/go-admin-team/go-admin-core/tools/database"
	. "github.com/go-admin-team/go-admin-core/tools/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"sync"
	"time"
	"warehouse/common/global"
	models2 "warehouse/common/models"
	"warehouse/models"
)

var once sync.Once
var db *gorm.DB

// Setup 配置数据库
func Setup() {
	for k := range toolsConfig.DatabasesConfig {
		setupSimpleDatabase(k, toolsConfig.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *toolsConfig.Database) {
	if global.Driver == "" {
		global.Driver = c.Driver
	}
	//log.Infof("%s => %s", host, pkg.Green(c.Source))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		log.Fatal(pkg.Red(c.Driver+" connect error :"), err)
	} else {
		log.Info(pkg.Green(c.Driver + " connect success !"))
	}

	e := mycasbin.Setup(db, "")

	sdk.Runtime.SetDb(host, db)
	sdk.Runtime.SetCasbin(host, e)
	err = db.Migrator().AutoMigrate(
		&models.MaterialSql{},
		&models.InboundDetailSql{},
		&models.OutboundDetailSql{},
		&models.InBoundPerson{},
		&models.OutBoundPerson{},
		&models.Unit{},
		&models.StorageLocation{},
		&models.MaterialID{},
		&models2.SysDept{},
		&models2.SysMenu{},
		&models2.SysRoleDept{},
		&models2.SysUser{},
		&models2.SysRole{},
		&models2.SysApi{},
	)
	if err != nil {
		log.Fatal("db.AutoMigrate failed ... ")
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
			err = db.Exec(line).Error
			if err != nil {
				log.Info("Error executing SQL statement: %s\n%s\n", err, line)
				continue
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	})
	//sqlStatements := make(map[string]string)
	//files := []string{"query1.sql", "query2.sql", "query3.sql"}
	//
	//for _, file := range files {
	//	content, err := ioutil.ReadFile(file)
	//	if err != nil {
	//		log.Fatalf("failed to read %s: %v", file, err)
	//	}
	//
	//	sqlStatements[file] = strings.TrimSpace(string(content))
	//}
	//query := sqlStatements["config/db.sql"]
	//rows, err := db.Raw(query).Rows()
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var col1 string
	//	var col2 int
	//	err := rows.Scan(&col1, &col2)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Printf("%s %d\n", col1, col2)
	//}
}
