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
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
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
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

// email
type Email struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
