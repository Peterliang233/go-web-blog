package model

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//添加文章
func CreateArticle(data *Article) int {
	err := databases.Db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}
//根据用户的id查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	if err := databases.Db.Where("id = ?", id).First(&article).Error; err != nil {
		return article, errmsg.ErrArticleNotExist
	}
	return article, errmsg.Success
}
//查询单个目录的id下面的所有文章,并且进行分页显示
func GetCategoryToArticles(id int, pageSize int, pageNum int) ([]Article, int, uint64) {
	var categoryArticleList []Article
	var total uint64
	databases.Db.Preload("category").Where("cid = ?", id).Find(&categoryArticleList).Count(total)
	err := databases.Db.Preload("category").
		Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).Find(&categoryArticleList).Error
	if err != nil {
		return nil, errmsg.ErrArticleNotExist, 0
	}
	return categoryArticleList, errmsg.Success, total
}
//查询文章列表
func GetArticles(PageSize, PageNum int) ([]Article, int) {
	var article []Article
	err := databases.Db.Limit(PageSize).Offset((PageNum-1)*PageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.Error
	}
	return article, errmsg.Success
}

//编辑文章
func EditArticle(id int, data *Article) int {
	var articleMap = make(map[string]interface{})
	articleMap["title"] = data.Title
	articleMap["cid"] = data.Cid
	articleMap["content"] = data.Content
	articleMap["desc"] = data.Desc
	articleMap["img"] = data.Img
	err := databases.Db.Table("article").Where("id = ?", id).
		Updates(articleMap).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//删除文章
func DelArticle(id int) int {
	var article Article
	if err := databases.Db.Where("id = ?", id).Delete(&article).Error; err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}