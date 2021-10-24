package user

import (
	"encoding/base64"
	"github.com/Peterliang233/go-blog/databases/mysql"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	_ "github.com/jordan-wright/email"
	"golang.org/x/crypto/scrypt"
	"log"
	_ "net/smtp"
)

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

// EditUser 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User

	var UserMap = make(map[string]interface{})

	UserMap["username"] = data.Username
	UserMap["id"] = data.ID
	err := mysql.Db.Model(&user).Where("id = ?", id).Updates(UserMap).Error

	if err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}

// CheckLogin 检查是否可以登录
func CheckLogin(username string, password string) int {
	var login model.User
	if err := mysql.Db.Where("username = ?", username).First(&login).Error; err != nil {
		return errmsg.Error
	}

	if ScryptPassword(password) != login.Password {
		return errmsg.ErrPassword
	}

	return errmsg.Success
}
