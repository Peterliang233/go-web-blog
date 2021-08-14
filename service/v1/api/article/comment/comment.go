package comment

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// CheckoutArticle 查询文章是否存在
func CheckoutArticle(id int) int {
	if err := databases.Db.Table("article").
		Where("id = ?", id).
		First(&model.Article{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errmsg.ErrArticleNotExist
		}

		return errmsg.Error
	}

	return errmsg.Success
}

// GetComments 通过文章的id获取所有评论
func GetComments(pageSize, pageNum, id int) ([]model.Comment, int, int) {
	var comments []model.Comment

	var total int

	if err := databases.Db.
		Select("author, id, content").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments).
		Error; err != nil {
		return nil, errmsg.Error, 0
	}

	if err := databases.Db.
		Table("comment").
		Where("aid = ?", id).Count(&total).
		Error; err != nil {
		return nil, errmsg.Error, 0
	}

	return comments, errmsg.Success, total
}

// CheckComment 检查评论
func CheckComment(id int) int {
	if err := databases.Db.
		Table("comment").
		Where("id = ?", id).First(&model.Comment{}).
		Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DelComment 删除评论
func DelComment(id int) int {
	if err := databases.Db.
		Where("id = ?", id).
		Delete(&model.Comment{}).
		Error; err != nil {

		return errmsg.Error
	}

	return errmsg.Success
}
