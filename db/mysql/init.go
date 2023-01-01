package mysql

import (
	"github.com/sjmshsh/IM/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

// MysqlDB 全局MysqlDB
var MysqlDB *gorm.DB

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		})
	// DSN 数据库连接池地址字符串拼接
	var builder strings.Builder
	s := []string{conf.MysqlUser, ":", conf.MysqlPassword, "@tcp(", conf.MysqlHost, ":", conf.MysqlPort, ")/", conf.MysqlName, "?charset=utf8mb4&parseTime=True&loc=Local"}
	for _, str := range s {
		builder.WriteString(str)
	}
	dsn := builder.String()
	// 数据库连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: newLogger,
	})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(conf.MysqlMaxIdleConns)
	sqlDB.SetConnMaxLifetime(conf.MysqlConnMaxLifetime)
	sqlDB.SetMaxOpenConns(conf.MysqlMaxOpenConns)
	if err != nil {
		log.Fatal(err)
	}
	MysqlDB = db
}
