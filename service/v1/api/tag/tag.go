package tag

import (
	"github.com/Peterliang233/go-blog/databases/mysql"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
)

// AddTag 添加标签
func AddTag(tag *model.Tag) (statusCode int, err error) {
	if err := mysql.Db.Create(tag).Error; err != nil {
		return errmsg.ErrTagCreate, err
	}

	return errmsg.Success, nil
}

// GetAllTags 获取所有的标签
func GetAllTags() ([]model.Tag, error) {
	var tags []model.Tag

	if err := mysql.Db.
		Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

// DeleteTag 删除标签
func DeleteTag(id int) (code int, err error) {
	if err := mysql.Db.
		Where("id = ?", id).
		Delete(&model.Tag{}).
		Error; err != nil {
		return errmsg.ErrTagDelete, err
	}

	return errmsg.Success, nil
}
