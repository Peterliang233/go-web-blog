package model

import "time"

// Article
type Article struct {
	ID         int        `json:"id"`
	TagID      int        `json:"tag_id"`
	CategoryID int        `json:"category_id"`
	Desc       string     `json:"desc"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreateAt   *time.Time `json:"create_at" gorm:"column:create_at"`
	DeleteAt   *time.Time `json:"delete_at" gorm:"column:delete_at"`
	UpdateAt   *time.Time `json:"update_at" gorm:"column:update_at"`
}

// User
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Like
type Like struct {
	ID        int    `json:"id"`
	ArticleID int    `json:"article_id"`
	Name      string `json:"name"`
}

// Tag
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment
type Comment struct {
	ID        int        `json:"id"`
	ArticleID int        `json:"article_id"`
	Content   string     `json:"content"`
	CreateAt  *time.Time `json:"create_at" gorm:"column:create_at"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:delete_at"`
}
