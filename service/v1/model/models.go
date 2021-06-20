package model

import "github.com/jinzhu/gorm"

// Article 文章-一个文章对应一个目录，一个目录对应多个文章
type Article struct {
	category Category `gorm:"foreignKey:cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(30);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(30);" json:"desc"`
	Content string `gorm:"type:text;not null" json:"content"`
	Uid     int    `gorm:"type:int;not null;" json:"uid"`
	Img     string `gorm:"type:varchar(30);" json:"img"`
}

// User 用户
type User struct {
	ID       int    `gorm:"type:int;not null;primaryKey;auto_increment" json:"id"`
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=12" label:"用户密码"`
	Email    string `gorm:"type:varchar(100);not null" json:"email" validate:"required,email" label:"邮箱"`
	Role     int    `gorm:"type:int; DEFAULT:2" json:"role" validate:"required,gte=2" label:"权限码"`
}

// Category 目录
type Category struct {
	ID   uint   `gorm:"type:int;not null;primary_key;auto_increment" json:"id"`
	Uid  int    `gorm:"type:int;not null;" json:"uid"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}

// Email 邮箱
type Email struct {
	ID        uint   `gorm:"type:int;not null;primary_key;auto_increment" json:"id"`
	EmailName string `gorm:"type:varchar(50);not null" json:"email_name" validate:"email" label:"邮箱"`
	Active    bool   `gorm:"type:tinyint(1);default:0" json:"active"`
}

// Comment 评论,多个评论对应一篇文章
type Comment struct {
	Article Article `gorm:"foreignKey:aid"`
	gorm.Model
	Aid     int    `gorm:"int;not null" json:"aid"`
	Content string `gorm:"text;not null" json:"content"`
	Author  string `gorm:"varchar(50);not null" json:"author"`
}
