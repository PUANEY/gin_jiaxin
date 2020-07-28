package users

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"shanshui/pkg/app"
	"shanshui/pkg/setting"
	"shanshui/pkg/utils"
	"time"
)

// jwt认证
type MyClaims struct {
	NickName string `json:"nickname"`
	jwt.StandardClaims
}
// jwt 过期时间
const TokenExpireDuration = time.Hour * 2
var MySecret = []byte("ShanShui")

// GenToken生成JWT
func GenToken(nickname string) (string, error) {
	m := MyClaims{
		nickname,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer: "ShanShuiZhuangShi",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析jwt
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/**
 * 登录
 * 检查用户名和密码
 */
func Login(c *gin.Context) {
	var (
		_   *User
		err error
	)
	g := app.Gin{C: c}
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")

	if nickname == "" || password == "" {
		g.Error("账号或密码不能为空")
	}

	// 判断用户是否存在
	_, err = GetUserByNickName(nickname)
	if err != nil {
		g.Error("该用户不存在")
	}

	_, err = CheckUserIsTrue(nickname, password)
	if err != nil {
		g.Error("Password error,please try again!")
	}
	// 登录成功后...
	tokenString, _ := GenToken(nickname)
	data := map[string]interface{}{
		"token": tokenString,
	}
	g.Success(data)
}

// 验证jwt
func AuthUser(c *gin.Context) {
	var user *User
	g := app.Gin{C: c}
	nickname := c.MustGet("nickname").(string)
	user, err := GetUserByNickName(nickname)
	if err != nil {
		g.Error("该账户不存在")
	}
	g.Success(map[string]interface{}{
		"nickname": nickname,
		"user": user,
	})
}

/**
 * 注册用户
 * 校验数据合法性
 */
func Register(c *gin.Context) {
	var u User
	g := app.Gin{C: c}
	if err := c.ShouldBind(&u); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			g.Error(err.Error())
		}
		// validator.ValidationErrors类型错误则进行翻译
		g.Success(utils.RemoveTopStruct(errs.Translate(utils.Trans)))
	}

	_, err := GetUserByNickName(u.NickName)
	// err == nil 说明用户名已存在
	if err == nil {
		g.Error("用户名已存在!")
		return
	}

	// 密码不一致情况
	if u.Password != u.RePassword {
		g.Error("密码不一致，请重新输入")
		return
	}
	CreateUser(u.NickName, u.Telephone, u.Password, u.RePassword, u.Age)
	g.Success("账户创建成功")
}

func UpAvatar(c *gin.Context)  {
	g := app.Gin{C:c}
	var user *User

	file, _ := c.FormFile("avatar")
	nickname := c.PostForm("nickname")

	// 对文件进行改名处理
	filename := utils.ChangeFileName(file.Filename)
	err := c.SaveUploadedFile(file, "./static/avatars/" + filename)
	if err != nil{
		g.Error("文件保存失败")
	}
	filepath := setting.ServerSetting.Host + setting.ServerSetting.HttpPort + "/static/avatars/" + filename

	// 执行更换头像操作
	user, err = GetUserByNickName(nickname)
	if err != nil {
		g.Error("修改失败")
	}
	// 找到用户后,执行更改操作
	user.AvatarUrl = filepath
	g.Success(map[string]interface{}{
		"user": user,
	})
}


