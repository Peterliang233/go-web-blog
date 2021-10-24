package like

import (
	"github.com/Peterliang233/go-blog/errmsg"
	likeService "github.com/Peterliang233/go-blog/service/v1/api/like"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Like 点赞接口
func Like(c *gin.Context) {
	email := c.Query("email")

	if ok, err := likeService.CheckLike(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"data":   email,
				"detail": errmsg.CodeMsg[errmsg.Error],
			},
		})
		return
	} else {
		// 点过赞了
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Success,
				"msg": map[string]interface{}{
					"data":   email,
					"detail": "已经点过赞了",
				},
			})
			return
		}

		// 放进缓存里面
		if err := likeService.RedisCreateEmail(email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": errmsg.ErrLikeCreate,
				"msg": map[string]interface{}{
					"data":   email,
					"detail": "点赞失败",
				},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.Success,
			"msg": map[string]interface{}{
				"data":   email,
				"detail": "点赞成功",
			},
		})
	}
}
