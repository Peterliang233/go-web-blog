package model

import "github.com/jinzhu/gorm"

//文章-一个文章对应一个目录，一个目录对应多个文章
type Article struct {
	category Category `gorm:"foreignKey:cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(30);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(30);" json:"desc"`
	Content string `gorm:"type:text;not null" json:"content"`
	Img     string `gorm:"type:varchar(30);" json:"img"`
}

//用户
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=12" label:"用户密码"`
	Role     int    `gorm:"type:int; DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限码"`
}

//目录
type Category struct {
	ID   uint   `gorm:"type:int;primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}
