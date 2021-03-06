package article

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/jinzhu/gorm"
)

//检查目录是否存在
func CheckCategory(data model.Category) int {
	var category model.Category
	err := databases.Db.Where("id = ?", data.ID).First(&category).Error
	if err == gorm.ErrRecordNotFound {
		err = databases.Db.Where("name = ?", data.Name).First(&category).Error
		if err == gorm.ErrRecordNotFound {
			return errmsg.Success
		} else if err != nil {
			return errmsg.ErrDatabaseNotFound
		} else {
			return errmsg.ErrCategoryIdUsed
		}
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	} else {
		return errmsg.ErrCategoryUsed
	}
}

//创建新的目录
func CreateCategory(data *model.Category) int {
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//获取分类的分页列表
func GetCategory(PageSize, PageNum int) ([]model.Category, int) {
	var category []model.Category
	err := databases.Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ErrCategoryNotExist
	}
	return category, errmsg.Success
}

//删除目录
func DeleteCategory(id int) int {
	var category model.Category
	err := databases.Db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//编辑目录,修改目录id对应的名字
func EditCategory(id int, name string) int {
	var category model.Category
	var UserMap = make(map[string]interface{})
	UserMap["name"] = name
	err := databases.Db.Model(&category).Where("id = ?", id).Updates(UserMap).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

func CheckCategoryName(name string) int {
	var category model.Category
	if err := databases.Db.Where("name = ?", name).First(&category).Error; err == gorm.ErrRecordNotFound {
		return errmsg.Success
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	}
	return errmsg.ErrCategoryUsed
}

func CheckCategoryId(id int) int {
	var category model.Category
	if err := databases.Db.Where("id = ?", id).First(&category).Error; err == gorm.ErrRecordNotFound {
		return errmsg.ErrCategoryNotExist
	} else if err != nil {
		return errmsg.ErrDatabaseNotFound
	}
	return errmsg.Success
}
