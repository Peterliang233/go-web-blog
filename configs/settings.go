package configs

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode    string
	HttpPort   string
	Secret     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	RdHost     string
	RdPort     string
	RdPassword string
	Addr       string
	Username   string
	Password   string
	Host       string
)

func init() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		fmt.Println("文件打开失败")
		os.Exit(1)
	}

	LoadData(cfg)

	LoadRedis(cfg)

	LoadServer(cfg)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
	Secret = file.Section("server").Key("Secret").MustString("secret")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("mysqlpassword")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadRedis(file *ini.File) {
	RdHost = file.Section("redis").Key("RdHost").MustString("localhost")
	RdPort = file.Section("redis").Key("RdPort").MustString("6379")
	RdPassword = file.Section("redis").Key("RdPassword").MustString("2562")
}
