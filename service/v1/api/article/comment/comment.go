package comment

import (
	"github.com/Peterliang233/go-blog/databases/mysql"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"net/http"
)

// CheckoutArticle 查询文章是否存在
func CheckoutArticle(id int) int {
	if err := mysql.Db.Table("article").
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

	if err := mysql.Db.
		Where("article_id = ?", id).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments).Count(&total).
		Error; err != nil {
		return nil, errmsg.Error, 0
	}

	return comments, errmsg.Success, total
}

// CheckComment 检查评论
func CheckComment(id int) int {
	if err := mysql.Db.
		Table("comment").
		Where("id = ?", id).First(&model.Comment{}).
		Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// DelComment 删除评论
func DelComment(id int) int {
	if err := mysql.Db.
		Where("id = ?", id).
		Delete(&model.Comment{}).
		Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// AddComment 添加评论
func AddComment(comment *model.Comment) int {
	if err := mysql.Db.Create(comment).Error; err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
