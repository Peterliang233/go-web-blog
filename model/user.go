package model

import (
	"encoding/base64"
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

//检查用户是否存在
func CheckUser(name string) int {
	var users User
	databases.Db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrUserNameUsed
	}
	return errmsg.Success
}

//创建新的用户
func CreateUser(data *User) int {
	data.Password = ScryptPassword(data.Password) //对密码进行加盐处理
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//获取用户分页列表
func GetUsers(PageSize, PageNum int) []User {
	var users []User
	err := databases.Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//密码加密（加盐操作）
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{23, 32, 21, 11, 11, 22, 11, 0}
	HashPassword, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(HashPassword)
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err := databases.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//编辑用户
func EditUser(id int, data *User) int {
	var user User
	var UserMap = make(map[string]interface{})
	UserMap["username"] = data.Username
	UserMap["id"] = data.ID
	err := databases.Db.Model(&user).Where("id = ?", id).Updates(UserMap).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

func CheckLogin(username string, password string) int {
	var login User
	if err := databases.Db.Where("username = ?", username).First(&login).Error; err != nil {
		return errmsg.Error
	}
	if ScryptPassword(password) != login.Password {
		return errmsg.ErrPassword
	}
	if login.Role != 0 {
		return errmsg.ErrNotHaveRight
	}
	return errmsg.Success
}
