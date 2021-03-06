package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-memo/conf"
	"go-memo/pkg/log"
)

var (
	DB                                                *gorm.DB
	err                                               error
	dbType, dbName, user, password, host, tablePrefix string
	c                                                 *conf.Config
)

func Migrate() {
	DB.AutoMigrate(
		&User{},
		&UserProfile{},
	)
}

func init() {
	c = conf.GetConfig()
	InitMysql(c.MySQL)
}

// 数据库初始化
func InitMysql(mysql conf.MySQLConfig) {
	dbType = mysql.Dbtype
	dbName = mysql.Dbname
	user = mysql.Username
	password = mysql.Password
	host = mysql.Host
	tablePrefix = mysql.Prefix

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName, ))

	if err != nil {
		log.Error(err.Error())
	}
	initDBConfig()
}

func InitSqlite() {
	var err error
	DB, err = gorm.Open("sqlite3", ":memory:")
	if conf.GetConfig().Common.Debug {
		DB.LogMode(true) // 开启 sql 日志
	}
	if err != nil {
		panic(err)
	}
	initDBConfig()
}

func initDBConfig() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if defaultTableName == tablePrefix+"casbin_rule" {
			return defaultTableName
		}
		return tablePrefix + defaultTableName
	}
	if c.Common.Debug {
		DB.LogMode(true) // 开启 sql 日志
	}
	DB.SingularTable(true) // 创建table名单数
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}


func CloseDB() {
	defer DB.Close()
}
