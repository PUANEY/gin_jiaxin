package users

import (
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine)  {
	// 登录
	e.POST("/login", Login)

	// 注册
	e.POST("/register", Register)

	// 验证token
	e.GET("/auth", JWTAuthMiddleware(), AuthUser)

	// 上传头像
	e.POST("/up_avatar", UpAvatar)
}
