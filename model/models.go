package model

// Article
type Article struct {
	ID         int    `gorm:"type:int;not null" json:"id"`
	CategoryId int    `gorm:"type:int;not null" json:"category_id"`
	TagId      int    `gorm:"type:int;not null" json:"tag_id"`
	Title      string `gorm:"type:varchar(30);not null" json:"title"`
	Desc       string `gorm:"type:varchar(30);" json:"desc"`
	Content    string `gorm:"type:text;not null" json:"content"`
}

// User
type User struct {
	ID       int    `gorm:"type:int;not null;primaryKey;auto_increment" json:"id"`
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=12" label:"用户密码"`
}

// Category
type Category struct {
	ID   uint   `gorm:"type:int;not null;primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(30);not null" json:"name"`
}

// Comment
type Comment struct {
	ID        int    `gorm:"int;not null" json:"id"`
	ArticleId int    `gorm:"int;not null" json:"article_id"`
	Content   string `gorm:"text;not null" json:"content"`
}
