package v1

import (
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加目录
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data)
	if code == errmsg.Success {
		code = model.CreateCategory(&data)
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
	data := model.GetCategory(page.PageSize, page.PageNum)
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
	code := model.DeleteCategory(id)
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
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.CheckCategory(category)
	if code == errmsg.Success {
		//执行更新的操作
		model.EditCategory(id, &category)
	}
	if code == errmsg.ErrCategoryUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"code": errmsg.CodeMsg[code],
			"data": category,
			"id":   id,
		},
	})
}
