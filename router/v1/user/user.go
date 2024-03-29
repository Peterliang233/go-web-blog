package user

import (
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	user2 "github.com/Peterliang233/go-blog/service/v1/api/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Page 查询目录
type Page struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

// EditUser 编辑用户的基本信息，但是不会包括修改密码
func EditUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(c.Param("id"))

	statusCode := http.StatusInternalServerError

	// 执行更新的操作
	code := user2.EditUser(id, &user)
	if code == errmsg.Success {
		statusCode = http.StatusOK
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
