package article

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/jinzhu/gorm"
)

func CheckoutArticle(id int) int {
	if err := databases.Db.Table("article").Where("id = ?", id).First(&model.Article{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ErrArticleNotExist
		} else {
			return errmsg.Error
		}
	} else {
		return errmsg.Success
	}
}

func GetComments(pageSize, pageNum, id int) ([]model.Comment, int, int) {
	var comments []model.Comment
	var total int
	if err := databases.Db.Select("author, id, content").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&comments).Error; err != nil {
		return nil, errmsg.Error, 0
	}
	if err := databases.Db.Table("comment").Where("aid = ?", id).Count(&total).Error; err != nil {
		return nil, errmsg.Error, 0
	}
	return comments, errmsg.Success, total
}

func CheckComment(id int) int {
	if err := databases.Db.Table("comment").Where("id = ?", id).First(&model.Comment{}).Error; err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

func DelComment(id int) int {
	if err := databases.Db.Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}
