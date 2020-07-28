package users

import (
	"github.com/gin-gonic/gin"
	"shanshui/pkg/app"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		g := app.Gin{C: c}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			g.Error("请求头中auth为空")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			g.Error("请求头中auth格式有误")
			return
		}

		// parts[1] 是获取到的tokenString，使用ParseToken函数解析

		mc, err := ParseToken(parts[1])
		if err != nil {
			g.Error("无效的Token")
			return
		}

		// 将当前请求的nickname信息保存到请求的上下文c上
		c.Set("nickname", mc.NickName)
		c.Next()
		// 后续的处理函数可以用c.Get("nickname")来获取当前请求的用户信息
	}
}
