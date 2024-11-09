package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
    
	"stu_Assistant/gateway/gwconfig"

)

// JWT token验证中间件, 
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			return
		}
		claims, err := gwconfig.ParseToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "权限过期，请重新登陆",
			})
			c.Abort()
			return
		}
		
		c.Set("user_id", claims.Id)  // 后续业务中可通过 c.Get("user_id") 获取用户id
		// ctl.InitUserInfo(c.Request.Context(), &ctl.UserInfo{Id: claims.Id})
		c.Next()
	}
}