package tag

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
)

// AddTag 添加标签
func AddTag(tag *model.Tag) (statusCode int, err error) {
	if err := databases.Db.Create(tag).Error; err != nil {
		return errmsg.ErrTagCreate, err
	}

	return errmsg.Success, nil
}
