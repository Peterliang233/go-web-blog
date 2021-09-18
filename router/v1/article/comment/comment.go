package comment

import (
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	"github.com/Peterliang233/go-blog/router/v1/user"
	commentService "github.com/Peterliang233/go-blog/service/v1/api/article/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var comment model.Comment
	_ = c.ShouldBindJSON(&comment)

	statusCode := http.StatusOK

	statusCode = commentService.AddComment(&comment)

	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg": map[string]interface{}{
			"data":   comment,
			"detail": errmsg.CodeMsg[statusCode],
		},
	})
}

// GetComment 获取相关的评论
func GetComment(c *gin.Context) {
	var page user.Page
	_ = c.ShouldBindJSON(&page)

	if page.PageSize == 0 {
		page.PageSize = -1
	}

	if page.PageNum == 0 {
		page.PageNum = -1
	}
	// 查询对应id文章的评论
	id, _ := strconv.Atoi(c.Param("id"))

	var code, total int

	var comments []model.Comment

	code = commentService.CheckoutArticle(id)

	statusCode := http.StatusOK

	if code != errmsg.Success {
		statusCode = http.StatusBadRequest
	}

	comments, code, total = commentService.GetComments(page.PageSize, page.PageNum, id)

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data":   comments,
			"total":  total,
			"detail": errmsg.CodeMsg[code],
		},
	})
}

// DelComment 删除相关评论
func DelComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := commentService.CheckComment(id)
	statusCode := http.StatusOK

	if code != errmsg.Success {
		statusCode = http.StatusNotFound
	}

	code = commentService.DelComment(id)

	if code != errmsg.Success {
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"data":   id,
			"detail": errmsg.CodeMsg[code],
		},
	})
}
