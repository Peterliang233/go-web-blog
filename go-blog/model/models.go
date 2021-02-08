package model

import "github.com/jinzhu/gorm"

//文章-一个文章对应一个目录，一个目录对应多个文章
type Article struct {
	category Category  `gorm:"foreignKey:cid"`
	gorm.Model
	Title string  `gorm:"type:varchar(30);not null" json:"title"`
	Cid int  `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(30);" json:"desc"`
	Content string  `gorm:"type:varchar(30);not null" json:"content"`
	Img string  `gorm:"type:varchar(30);" json:"img"`
}
//用户
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role int `gorm:"type:int" json:"role"`
}
//目录
type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}