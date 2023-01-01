package conf

import (
	"github.com/go-ini/ini"
	"log"
	"strings"
	"time"
)

var (
	HttpPort string

	MysqlHost            string
	MysqlPort            string
	MysqlUser            string
	MysqlPassword        string
	MysqlName            string
	MysqlIsLog           bool
	MysqlIsSingularTable bool
	MysqlMaxIdleConns    int
	MysqlMaxOpenConns    int
	MysqlConnMaxLifetime time.Duration

	RedisAddr     string
	RedisPassword string
	RedisDbName   int

	MangoAuthMechanism  string
	MangoUser           string
	MangoPassword       string
	MangoHosts          []string
	MangoDBName         string
	MangoConnectTimeout time.Duration
	MangoMaxPoolSize    uint64
	MangoMinPoolSize    uint64
)

func Init(path string) {
	file, err := ini.Load(path)
	if err != nil {
		log.Fatalln("Fail to parse 'conf/app.ini': ", err)
	}

	loadService(file)
	loadMysql(file)
	loadRedis(file)
	loadMongo(file)
}

func loadService(file *ini.File) {
	HttpPort = file.Section("service").Key("HttpPort").MustString(":8080")
}

func loadMysql(file *ini.File) {
	section, err := file.GetSection("mysql")
	if err != nil {
		log.Fatalln(err)
	}
	MysqlHost = section.Key("MysqlHost").String()
	MysqlPort = section.Key("MysqlPort").String()
	MysqlUser = section.Key("MysqlUser").String()
	MysqlPassword = section.Key("MysqlPassword").String()
	MysqlName = section.Key("MysqlName").String()
	MysqlIsLog = section.Key("MysqlIsLog").MustBool(true)
	MysqlIsSingularTable = section.Key("MysqlIsSingularTable").MustBool(true)
	MysqlMaxIdleConns = section.Key("MysqlMaxIdleConns").MustInt(20)
	MysqlMaxOpenConns = section.Key("MysqlMaxOpenConns").MustInt(100)
	MysqlConnMaxLifetime = time.Duration(section.Key("MysqlConnMaxLifetime").MustInt(30)) * time.Second
}

func loadRedis(file *ini.File) {
	section, err := file.GetSection("redis")
	if err != nil {
		log.Fatalln(err)
	}
	RedisAddr = section.Key("RedisAddr").String()
	RedisPassword = section.Key("RedisPassword").MustString("")
	RedisDbName = section.Key("RedisDbName").MustInt(1) // MustInt，defaultVal为1
}

func loadMongo(file *ini.File) {
	section, err := file.GetSection("mongo")
	if err != nil {
		log.Fatalln(err)
	}
	MangoAuthMechanism = section.Key("MangoAuthMechanism").String()
	MangoUser = section.Key("MangoUser").String()
	MangoPassword = section.Key("MangoPassword").String()
	// MangoHosts比较特殊，需要一个[]string。所以在ini文件中以,进行字符串分割
	MangoHosts = strings.Split(section.Key("MangoHosts").String(), ",")
	MangoDBName = section.Key("MangoDBName").String()
	MangoConnectTimeout = time.Duration(section.Key("MangoConnectTimeout").MustInt(10)) * time.Second
	MangoMaxPoolSize = section.Key("MangoMaxPoolSize").MustUint64(20)
	MangoMinPoolSize = section.Key("MangoMinPoolSize").MustUint64(5)
}