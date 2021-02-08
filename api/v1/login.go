package v1

import (
	"github.com/Peterliang233/go-blog/middleware"
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/utils"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)


type user struct {
	Username string  `json:"username"`
	Password string `json:"password"`
}
func AuthHandler(c *gin.Context) {
	var user user
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code" : errmsg.ErrRequest,
			"msg" : map[string]interface{}{
				"detail": "无效的参数",
				"data" : "",
			},
		})
		return
	}
	//超级管理员，初始化配置文件
	if user.Username == utils.Username && user.Password == utils.Password {
		tokenString, code := middleware.GenerateToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg": map[string]interface{}{
				"token" : tokenString,
				"code" : errmsg.CodeMsg[code],
				"detail" : "登录成功",
			},
		})
	}else {
		//检查是否具有登录权限
		code := model.CheckLogin(user.Username, user.Password)
		c.JSON(http.StatusOK, gin.H{
			"status" : code,
			"msg": map[string]interface{}{
				"code" : errmsg.CodeMsg[code],
			},
		})
	}
}