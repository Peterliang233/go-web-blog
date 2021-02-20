package email

import (
	"fmt"
	"github.com/Peterliang233/go-blog/configs"
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/http"
	"net/smtp"
	"strconv"
	"time"
)

//发送邮件
func SendEmail(Email, username string) int {
	e := email.NewEmail()
	e.From = "Peterliang <ncuyanping666@126.com>"
	e.To = []string{Email}
	e.Subject = "博客注册验证通知"
	e.Text = []byte("博客注册验证通知\n您好！\n" + username + "<" + Email + ">,您的邮箱在" +
		time.Now().Format("2006-01-02 15:04:05") + "被用于注册ginBlog.\n 请您打开以下链接进行邮箱验证：\n" +
		fmt.Sprintf("%s%s/api/v1/email/:1",
			configs.DbHost,
			configs.HttpPort) +
		"\n感谢您的使用，希望您使用愉快^_^\nginBlog官方团队:Peterliang")
	err := e.Send(configs.Addr, smtp.PlainAuth("", configs.Username, configs.Password, configs.Host))
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

//验证邮箱
func VerifyEmail(c *gin.Context) {
	var Email model.Email
	id, err := strconv.Atoi(c.Param("id"))
	Email.EmailName = c.Query("email")
	var code int
	statusCode := http.StatusOK
	if err != nil {
		code = errmsg.ErrRequest
		statusCode = http.StatusBadRequest
	}
	if id == 1 {
		code = errmsg.Success
		Email.Active = true //激活该邮箱
		if err = databases.Db.Create(&Email).Error; err != nil {
			statusCode = http.StatusInternalServerError
		}
	} else {
		code = errmsg.Error
	}
	c.JSON(statusCode, gin.H{
		"code": code,
		"data": map[string]interface{}{
			"email": Email.EmailName,
			"msg":   errmsg.CodeMsg[code],
		},
	})
}
