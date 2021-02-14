package v1

import (
	"github.com/Peterliang233/go-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

//上传数据接口
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrRequest,
			"msg": map[string]interface{}{
				"status": errmsg.CodeMsg[errmsg.ErrRequest],
			},
		})
	} else {
		dist := path.Join("./", file.Filename)
		_ = c.SaveUploadedFile(file, dist)
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.Success,
			"msg": map[string]interface{}{
				"data":   "",
				"detail": "upload success",
			},
		})
	}
}
