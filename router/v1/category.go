package v1

import (
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/api/article"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加目录
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := article.CheckCategory(data)
	if code == errmsg.Success {
		code = article.CreateCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data":   data,
			"status": errmsg.CodeMsg[code],
		},
	})
}

//查询目录
func GetCategory(c *gin.Context) {
	var page Page
	_ = c.ShouldBindJSON(&page)
	if page.PageSize == 0 {
		page.PageSize = -1
	}
	if page.PageNum == 0 {
		page.PageNum = -1
	}
	data := article.GetCategory(page.PageSize, page.PageNum)
	code := errmsg.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"status": errmsg.CodeMsg[code],
			"data":   data,
		},
	})
}

//删除目录
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := article.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"id":   id,
			"code": errmsg.CodeMsg[code],
		},
	})
}

//编辑目录的基本信息
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Query("name")
	code := article.CheckCategoryName(name)
	if code == errmsg.Success {
		//执行更新的操作
		code = article.CheckCategoryId(id)
		if code == errmsg.Success {
			article.EditCategory(id, name)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"code": errmsg.CodeMsg[code],
			"data": name,
			"id":   id,
		},
	})
}
