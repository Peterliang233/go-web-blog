package utils

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
	Username   string
	Password   string
)

func init() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		fmt.Println("文件打开失败")
		os.Exit(1)
	}
	LoadData(cfg)
	LoadServer(cfg)
	LoadLogin(cfg)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("AppMode")
	HttpPort = file.Section("server").Key("HttpPort").MustString("HttpPort")
	Secret = file.Section("server").Key("Secret").MustString("Secret")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost  = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("mysqlpassword")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadLogin(file *ini.File) {
	Username = file.Section("login").Key("Username").MustString("Peterliang")
	Password = file.Section("login").Key("Password").MustString("666")
}