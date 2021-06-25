package user

import (
	"encoding/base64"
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jordan-wright/email"
	"golang.org/x/crypto/scrypt"
	"log"
	_ "net/smtp"
)

//检查用户名和邮箱是否存在
func CheckUser(name string, email string) int {
	var users model.User
	if err := databases.Db.Table("user").Where("username = ?", name).First(&users).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return errmsg.ErrDatabaseNotFound
		}
		if err = databases.Db.Table("user").Where("email = ?", email).First(&users).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return errmsg.ErrDatabaseNotFound
			} else {
				return errmsg.Success
			}
		} else {
			return errmsg.ErrUserEmailUsed //用户邮箱存在
		}
	} else {
		return errmsg.ErrUserNameUsed //用户名存在
	}
}

//创建新的用户
func CreateUser(data *model.User) int {
	data.Password = ScryptPassword(data.Password) //对密码进行加盐处理
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//获取用户分页列表
func GetUsers(PageSize, PageNum int) ([]model.User, int, int) {
	var users []model.User
	var total int
	var code int
	err := databases.Db.Select("id,username,role").Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			code = errmsg.ErrUserNotExist
		} else {
			code = errmsg.Error
		}
		return nil, 0, code
	}
	err = databases.Db.Table("user").Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.Error
	}
	return users, total, errmsg.Success
}

// ScryptPassword 密码加密（加盐操作）
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{23, 32, 21, 11, 11, 22, 11, 0}
	HashPassword, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	//以上是加盐的一些过程
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(HashPassword)
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user model.User
	err := databases.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// EditUser 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User
	var UserMap = make(map[string]interface{})
	UserMap["username"] = data.Username
	UserMap["email"] = data.Email
	UserMap["id"] = data.ID
	err := databases.Db.Model(&user).Where("id = ?", id).Updates(UserMap).Error
	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

func CheckLogin(username string, password string) int {
	var login model.User
	if err := databases.Db.Where("username = ?", username).First(&login).Error; err != nil {
		return errmsg.Error
	}

	if ScryptPassword(password) != login.Password {
		return errmsg.ErrPassword
	}

	if login.Role > 2 {
		return errmsg.ErrNotHaveRight
	}

	return errmsg.Success
}

func GetRight(username string) (code int) {
	var user model.User
	if err := databases.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return errmsg.Error
	}

	if user.Role != 1 {
		return errmsg.ErrUserNotHaveAddRight
	}

	return errmsg.Success
}
