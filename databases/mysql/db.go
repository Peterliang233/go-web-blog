package mysql

import (
	"fmt"
	"github.com/Peterliang233/go-blog/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var Db *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	var err error
	Db, err = gorm.Open(configs.Db,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			configs.DbUser,
			configs.DbPassword,
			configs.DbHost,
			configs.DbPort,
			configs.DbName,
		))

	Db.SingularTable(true)

	if err != nil {
		fmt.Println("数据库打开失败")
	}

	Db.DB().SetMaxIdleConns(10)

	Db.DB().SetMaxOpenConns(100)

	Db.DB().SetConnMaxLifetime(10 * time.Second)
}
