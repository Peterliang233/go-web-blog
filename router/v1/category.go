package v1

import (
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/service/v1/api/article"
	"github.com/Peterliang233/go-blog/service/v1/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加目录
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	statusCode := http.StatusOK
	code := article.CheckCategory(data)

	if code == errmsg.Success {
		code = article.CreateCategory(&data)
		if code != errmsg.Success {
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusNotFound
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data":   data,
			"status": errmsg.CodeMsg[code],
		},
	})
}

// GetCategory 查询目录
func GetCategory(c *gin.Context) {
	var page Page
	_ = c.ShouldBindJSON(&page)

	if page.PageSize == 0 {
		page.PageSize = -1
	}

	if page.PageNum == 0 {
		page.PageNum = -1
	}

	statusCode := http.StatusOK
	data, code := article.GetCategory(page.PageSize, page.PageNum)
	if code != errmsg.Success {
		statusCode = http.StatusNotFound
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"status": errmsg.CodeMsg[code],
			"data":   data,
		},
	})
}

// DelCategory 删除目录
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := article.DeleteCategory(id)
	statusCode := http.StatusOK
	if code != errmsg.Success {
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

// EditCategory 编辑目录的基本信息
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Query("name")
	statusCode := http.StatusOK
	code := article.CheckCategoryName(name)
	if code != errmsg.Success {
		statusCode = http.StatusBadRequest
	} else {
		//执行更新的操作
		code = article.CheckCategoryId(id)
		if code != errmsg.Success {
			statusCode = http.StatusBadRequest
			code = article.EditCategory(id, name)
			if code != errmsg.Success {
				statusCode = http.StatusInternalServerError
			}
		}
	}
	c.JSON(statusCode, gin.H{
		"status": code,
		"msg": map[string]interface{}{
			"code": errmsg.CodeMsg[code],
			"data": name,
			"id":   id,
		},
	})
}
