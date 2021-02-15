package model

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//检查目录是否存在
func CheckCategory(data Category) int {
	var category Category
	err := databases.Db.Where("name = ?", data.Name).First(&category).Error
	if err == gorm.ErrRecordNotFound {
		err = databases.Db.Where("id = ?", data.ID).First(&category).Error
		if err == gorm.ErrRecordNotFound {
			return errmsg.Success
		} else if err != nil {
			return errmsg.ErrDatabaseFind
		} else {
			return errmsg.ErrCategoryIdUsed
		}
	} else if err != nil {
		return errmsg.ErrDatabaseFind
	} else {
		return errmsg.ErrCategoryUsed
	}
}

//创建新的目录
func CreateCategory(data *Category) int {
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//获取用户分页列表
func GetCategory(PageSize, PageNum int) []Category {
	var category []Category
	err := databases.Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return category
}

//删除目录
func DeleteCategory(id int) int {
	var category Category
	err := databases.Db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//编辑目录
func EditCategory(id int, data *Category) int {
	var category Category
	var UserMap = make(map[string]interface{})
	UserMap["name"] = data.Name
	err := databases.Db.Model(&category).Where("id = ?", id).Updates(UserMap).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}
