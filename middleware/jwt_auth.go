package middleware

import (
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

//验证是否登录
func JwtAuth() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		//如果在redis缓存中找到token的话，说明该token已注销
		ok := redis_service.CheckLogout(token)
		if ok {
			res.FailWithMsg("token已注销", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}



// // 验证是否登录
// func JwtAuth(c *gin.Context) {
// 	token := c.Request.Header.Get("token")
// 	if token == "" {
// 		res.FailWithMsg("未携带token", c)
// 		c.Abort()
// 		return
// 	}
// 	claims, err := jwts.ParseToken(token)
// 	if err != nil {
// 		res.FailWithMsg("token错误", c)
// 		c.Abort()
// 		return
// 	}
// 	c.Set("claims", claims)
// }
