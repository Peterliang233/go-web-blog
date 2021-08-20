package category

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// CheckCategory 检查分类是否存在
func CheckCategory(data model.Category) int {
	var category model.Category
	err := databases.Db.Where("id = ?", data.ID).First(&category).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = databases.Db.Where("name = ?", data.Name).First(&category).Error

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return errmsg.Success
		case err != nil:
			return errmsg.ErrDatabaseNotFound
		default:
			return errmsg.ErrCategoryIDUsed
		}
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	}

	return errmsg.ErrCategoryUsed
}

// CreateCategory 创建新的分类
func CreateCategory(data *model.Category) int {
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// GetCategory 获取分类的分页列表
func GetCategory(pageSize, pageNum int) ([]model.Category, int) {
	var category []model.Category
	err := databases.Db.
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&category).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errmsg.ErrCategoryNotExist
	}

	return category, errmsg.Success
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category model.Category
	if err := databases.Db.
		Where("id = ?", id).
		Delete(&category).
		Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// EditCategory 编辑分类,修改目录id对应的名字
func EditCategory(id int, name string) int {
	var category model.Category

	var UserMap = make(map[string]interface{})

	UserMap["name"] = name
	if err := databases.Db.
		Model(&category).
		Where("id = ?", id).
		Updates(UserMap).
		Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// CheckCategoryName 检查该分类是否存在
func CheckCategoryName(name string) int {
	var category model.Category
	if err := databases.Db.
		Where("name = ?", name).
		First(&category).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.Success
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	}

	return errmsg.ErrCategoryUsed
}

// CheckCategoryID 检查分类的id是否存在
func CheckCategoryID(id int) int {
	var category model.Category
	if err := databases.Db.
		Where("id = ?", id).
		First(&category).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.ErrCategoryNotExist
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	}

	return errmsg.Success
}
