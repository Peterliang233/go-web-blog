package model

import (
	"encoding/base64"
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	_ "github.com/jordan-wright/email"
	"golang.org/x/crypto/scrypt"
	"log"
	"net/smtp"
	_ "net/smtp"
	"time"
)

//检查用户名和邮箱是否存在
func CheckUser(name string, email string) int {
	var users User
	if err := databases.Db.Table("user").Where("username = ?", name).First(&users).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return errmsg.ErrDatabaseFind
		}
		if err = databases.Db.Table("user").Where("email = ?", email).First(&users).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return errmsg.ErrDatabaseFind
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
func CreateUser(data *User) int {
	data.Password = ScryptPassword(data.Password) //对密码进行加盐处理
	err := databases.Db.Create(data).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//获取用户分页列表
func GetUsers(PageSize, PageNum int) ([]User, int) {
	var users []User
	var total int
	err := databases.Db.Select("id,username,role").Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil {
		return nil, 0
	}
	err = databases.Db.Table("user").Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return users, total
}

//密码加密（加盐操作）
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
	UserMap["email"] = data.Email
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
	if login.Role > 2 {
		return errmsg.ErrNotHaveRight
	}
	return errmsg.Success
}

func GetRight(username string) (code int) {
	var user User
	if err := databases.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return errmsg.Error
	}
	if user.Role != 1 {
		return errmsg.ErrUserNotHaveAddRight
	}
	return errmsg.Success
}

func SendEmail(Email string) int {
	e := email.NewEmail()
	e.From = "Peterliang <ncuyanping666@126.com>"
	e.To = []string{Email}
	e.Subject = "博客注册通知"
	e.Text = []byte("博客注册通知！！！\n亲爱的" + Email + ",您的邮箱在" +
		time.Now().Format("2006-01-02 15:04:05") + "被用于注册ginBlog，感谢您的使用，希望您使用愉快^_^\nPeterliang")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "ncuyanping666@126.com", "OICRHJRGCHSPAAIZ", "smtp.126.com"))
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}
