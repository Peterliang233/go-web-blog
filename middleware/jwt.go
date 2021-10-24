package middleware

import (
	"github.com/Peterliang233/go-blog/configs"
	"github.com/Peterliang233/go-blog/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var MySecret = []byte(configs.Secret)

type MyClaims struct {
	Username string `json:"username"` // 利用中间件保存一些有用的信息
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, int) {
	Claims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // 设置过期时间
			Issuer:    "peter",                              // 设置签发人
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	token, err := reqClaims.SignedString(MySecret)

	if err != nil {
		return "", errmsg.Error
	}

	return token, errmsg.Success
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// JWTAuthMiddleware jwt中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == " " {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.CodeMsg[errmsg.AuthEmpty],
					"msg":    "请求头中的auth格式有误",
				},
			})
			c.Abort()

			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.InvalidToken,
					"msg":    errmsg.CodeMsg[errmsg.InvalidToken],
				},
			})
			c.Abort()

			return
		}

		claims, err := ParseToken(parts[1])

		// token失效
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.InvalidToken,
					"msg":    err.Error(),
				},
			})
			c.Abort()

			return
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
