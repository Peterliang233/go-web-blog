package v1

import (
	"github.com/Peterliang233/go-blog/errmsg"
	user2 "github.com/Peterliang233/go-blog/service/v1/api/user"
	"github.com/Peterliang233/go-blog/service/v1/api/user/validator"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	username := c.MustGet("username").(string)
	code := user2.GetRight(username)
	if code != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"status": errmsg.CodeMsg[code],
				"data":   "",
			},
		})
		return
	}
	msg, code := validator.Validate(&data)
	//进行数据的验证
	if code != errmsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": msg,
				"status": errmsg.CodeMsg[code],
				"data":   "",
			},
		})
		return
	}
	code = user2.CheckUser(data.Username, data.Email) //检查用户名和邮箱是否已经被使用
	if code == errmsg.Success {
		code = user2.CreateUser(&data)
		if code == errmsg.Success {
			user2.SendEmail(data.Email)
		}
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

//查询目录
type Page struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func GetUsers(c *gin.Context) {
	var page Page
	_ = c.ShouldBindJSON(&page)
	if page.PageSize == 0 {
		page.PageSize = -1
	}
	if page.PageNum == 0 {
		page.PageNum = -1
	}
	data, total := user2.GetUsers(page.PageSize, page.PageNum)
	code := errmsg.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"status": errmsg.CodeMsg[code],
			"data":   data,
			"total":  total,
		},
	})
}

//删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := user2.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"id":   id,
			"code": errmsg.CodeMsg[code],
		},
	})
}

//编辑用户的基本信息，但是不会包括修改密码
func EditUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	code := user2.CheckUser(user.Username, user.Email)
	if code == errmsg.Success {
		//执行更新的操作
		user2.EditUser(id, &user)
	}
	//if code == errmsg.ErrUserNameUsed {
	//	c.Abort()
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"code": errmsg.CodeMsg[code],
			"data": user,
			"id":   id,
		},
	})
}
