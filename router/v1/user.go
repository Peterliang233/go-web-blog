package v1

import (
	"github.com/Peterliang233/go-blog/databases"
	"github.com/Peterliang233/go-blog/errmsg"
	user2 "github.com/Peterliang233/go-blog/service/v1/api/user"
	"github.com/Peterliang233/go-blog/service/v1/api/user/email"
	"github.com/Peterliang233/go-blog/service/v1/api/user/validator"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// VerifyUser 验证添加注册用户的信息
func VerifyUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	username := c.MustGet("username").(string)
	code := user2.GetRight(username)

	if code != errmsg.Success {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"status": errmsg.CodeMsg[code],
				"data":   "",
			},
		})

		return
	}
	msg, code := validator.Validate(&data)
	// 进行数据的验证
	if code != errmsg.Success {
		c.JSON(http.StatusNotFound, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": msg,
				"status": errmsg.CodeMsg[code],
				"data":   "",
			},
		})

		return
	}

	code = user2.CheckUser(data.Username, data.Email)

	if code == errmsg.Success {
		email.SendEmail(data.Email, data.Username)
	}

	c.JSON(http.StatusOK, gin.H{

		"code": code,
		"msg": map[string]interface{}{
			"data": map[string]string{
				"username": data.Username,
				"email":    data.Email,
			},
			"status": errmsg.CodeMsg[code],
		},
	})
}

// Register 注册账户接口
func Register(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code, statusCode int
	if err := databases.Db.Table("email").Where("email_name = ?", data.Email).First(&model.Email{}).Error; err != nil {
		code = errmsg.ErrEmailUnVerify
		statusCode = http.StatusForbidden
	} else {
		code = user2.CreateUser(&data)
		statusCode = http.StatusOK
	}
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data": map[string]string{
				"username": data.Username,
				"email":    data.Email,
			},
			"status": errmsg.CodeMsg[code],
		},
	})
}

// Page 查询目录
type Page struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

// GetUsers 获取所有的用户
func GetUsers(c *gin.Context) {
	var page Page
	var statusCode int
	_ = c.ShouldBindJSON(&page)
	if page.PageSize == 0 {
		page.PageSize = -1
	}
	if page.PageNum == 0 {
		page.PageNum = -1
	}
	data, total, code := user2.GetUsers(page.PageSize, page.PageNum)
	if code == errmsg.Success {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"status": errmsg.CodeMsg[code],
			"data":   data,
			"total":  total,
		},
	})
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := user2.DeleteUser(id)
	var statusCode int
	if code == errmsg.Success {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	c.JSON(statusCode, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"id":   id,
			"code": errmsg.CodeMsg[code],
		},
	})
}

// EditUser 编辑用户的基本信息，但是不会包括修改密码
func EditUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	code := user2.CheckUser(user.Username, user.Email)
	statusCode := http.StatusInternalServerError
	if code == errmsg.Success {
		//执行更新的操作
		code = user2.EditUser(id, &user)
		if code == errmsg.Success {
			statusCode = http.StatusOK
		}
	}
	c.JSON(statusCode, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"code": errmsg.CodeMsg[code],
			"data": user,
			"id":   id,
		},
	})
}
