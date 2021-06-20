package article

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/jinzhu/gorm"
)

// CreateArticle 添加文章
func CreateArticle(data *model.Article) int {
	err := databases.Db.Create(&data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetArticle 根据用户的id查询单个文章
func GetArticle(id int) (model.Article, int) {
	var article model.Article
	if err := databases.Db.Where("id = ?", id).First(&article).Error; err != nil {
		return article, errmsg.ErrArticleNotExist
	}

	return article, errmsg.Success
}

// GetCategoryToArticles 查询单个目录的id下面的所有文章,并且进行分页显示
func GetCategoryToArticles(id int, pageSize int, pageNum int) ([]model.Article, int, uint64) {
	var categoryArticleList []model.Article
	var total uint64
	if err := databases.Db.Table("article").Where("cid = ?", id).Find(&categoryArticleList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errmsg.ErrArticleNotExist, 0
		}
		return nil, errmsg.ErrDatabaseNotFound, 0
	}
	err := databases.Db.Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).Find(&categoryArticleList).Error
	if err != nil {
		return nil, errmsg.ErrDatabaseNotFound, 0
	}
	if err := databases.Db.Table("article").Where("cid = ?", id).Count(&total).Error; err != nil {
		return nil, errmsg.ErrDatabaseNotFound, 0
	}
	return categoryArticleList, errmsg.Success, total
}

// GetArticles 查询文章列表
func GetArticles(PageSize, PageNum int) ([]model.Article, int) {
	var article []model.Article
	err := databases.Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.Error
	}
	return article, errmsg.Success
}

// EditArticle 编辑文章
func EditArticle(id int, data *model.Article) int {
	var articleMap = make(map[string]interface{})
	articleMap["title"] = data.Title
	articleMap["cid"] = data.Cid
	articleMap["content"] = data.Content
	articleMap["desc"] = data.Desc
	articleMap["img"] = data.Img
	err := databases.Db.Table("article").Where("id = ?", id).
		Updates(articleMap).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ErrArticleNotExist
		} else {
			return errmsg.Error
		}
	}
	return errmsg.Success
}

// DelArticle 删除文章
func DelArticle(id int) int {
	var article model.Article
	if err := databases.Db.Where("id = ?", id).Delete(&article).Error; err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}
