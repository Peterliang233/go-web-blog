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
	LoadServer(cfg)
	LoadEmail(cfg)
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

func LoadEmail(file *ini.File) {
	Addr = file.Section("email").Key("Addr").MustString("smtp.126.com:25")
	Username = file.Section("email").Key("Username").MustString("ncuyanping666@126.com")
	Password = file.Section("email").Key("Password").MustString("OICRHJRGCHSPAAIZ")
	Host = file.Section("email").Key("Host").MustString("smtp.126.com")
}
