package model

import (
	"fmt"
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/utils"
	"github.com/jinzhu/gorm"
	"time"
)

//初始化数据库
func InitDb() {
	var err error
	databases.Db, err = gorm.Open(utils.Db,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			utils.DbUser,
			utils.DbPassword,
			utils.DbHost,
			utils.DbPort,
			utils.DbName,
		))
	databases.Db.SingularTable(true) //不给表的名字加复数
	databases.Db.AutoMigrate(&Article{})
	//databases.Db.AutoMigrate(&User{})
	databases.Db.AutoMigrate(&Category{})
	if err != nil {
		fmt.Println("数据库打开失败")
	}
	databases.Db.DB().SetMaxIdleConns(10)
	databases.Db.DB().SetMaxOpenConns(100)
	databases.Db.DB().SetConnMaxLifetime(10 * time.Second)
	//err = Db.Close()
	//if err != nil {
	//	fmt.Println("数据库关闭失败")
	//}
}
