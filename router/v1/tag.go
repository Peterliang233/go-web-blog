package v1

import (
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/Peterliang233/go-blog/model"
	TagService "github.com/Peterliang233/go-blog/service/v1/api/tag"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	var tag model.Tag
	_ = c.ShouldBind(&tag)

	statusCode, err := TagService.AddTag(&tag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": statusCode,
			"msg": map[string]interface{}{
				"data":   tag,
				"detail": errmsg.CodeMsg[statusCode],
			},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"msg": map[string]interface{}{
			"data":   tag,
			"detail": errmsg.CodeMsg[statusCode],
		},
	})
}

// GetAllTags 获取所有的标签
func GetAllTags(c *gin.Context) {
	tags, err := TagService.GetAllTags()

	statusCode := http.StatusOK

	code := errmsg.Success

	if err != nil {
		statusCode = http.StatusInternalServerError
		code = errmsg.ErrTagGet
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   tags,
		},
	})
}

// DelTag 删除标签
func DelTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
				"data":   id,
			},
		})

		return
	}

	code, err := TagService.DeleteTag(id)

	statusCode := http.StatusOK

	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   id,
		},
	})
}
