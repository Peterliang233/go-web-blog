package v1

import (
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"msg" : map[string]interface{} {
			"data": data,
			"code" : errmsg.CodeMsg[code],
		},
	})
}

//查询文章列表
func GetArticles(c *gin.Context) {
	var page Page
	_ = c.ShouldBindJSON(&page)
	if page.PageSize == 0 {
		page.PageSize = -1
	}
	if page.PageNum == 0 {
		page.PageNum = -1
	}
	data, code := model.GetArticles(page.PageSize, page.PageNum)
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"msg": map[string]interface{}{
			"code" : errmsg.CodeMsg[code],
			"data" : data,
		},
	})
}

//根据文章的id查找对应的文章
func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : errmsg.ErrRequest,
			"msg": map[string]interface{}{
				"data": "",
				"status": errmsg.CodeMsg[errmsg.ErrRequest],
			},
		})
	}
	article, code := model.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg" : map[string]interface{}{
			"data" : article,
			"code" : errmsg.CodeMsg[code],
		},
	})
}
//查询某一个目录下面的所有的文章
func GetCategoryToArticle(c *gin.Context) {
	var page Page
	_ = c.ShouldBindJSON(&page)
	id, _ := strconv.Atoi(c.Param("id"))
	if page.PageSize == 0 {
		page.PageSize = -1
	}
	if page.PageNum == 0 {
		page.PageNum = -1
	}
	data, code, total := model.GetCategoryToArticles(id, page.PageSize, page.PageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg" : map[string]interface{}{
			"data": data,
			"total": total,
			"code": errmsg.CodeMsg[code],
		},
	})
}

//删除文章
func DelArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DelArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"msg": map[string]interface{}{
			"id" : id,
			"code": errmsg.CodeMsg[code],
		},
	})
}



//编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.EditArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"msg" : map[string]interface{}{
			"code" : errmsg.CodeMsg[code],
			"data" : article,
			"id" : id,
		},
	})
}