package users

import (
	"github.com/jinzhu/gorm"
	"shanshui/common"
)

// 用户表
type User struct {
	gorm.Model
	NickName string	`gorm:"not null;unique" json:"nickname" binding:"required" form:"nickname"`
	AvatarUrl string `gorm:"default:null" json:"avatar_url" form:"avatar_url"`
	Age uint `gorm:"not null" json:"age" binding:"gte=1,lte=130,required" form:"age"`
	Telephone string `gorm:"type:varchar(20)" json:"telephone" binding:"required" form:"telephone"`
	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required" form:"password"`
	RePassword string `gorm:"type:varchar(20);not null" json:"re_password" binding:"required" form:"re_password"`
}

// 设置User的表名为`users`
func (User) TableName() string {
	return "users"
}

func GetUserByNickName(nickname string) (*User, error)  {
	var user User
	err := common.DB.Where("nick_name = ?", nickname).First(&user).Error
	return &user, err
}

func CheckUserIsTrue(nickname, password string)(*User, error) {
	var user User
	err := common.DB.Where(&User{NickName: nickname, Password: password}).First(&user).Error
	return &user, err
}

func CreateUser(nickname, telephone, password, rePassword string, age uint) {
	user := User{NickName: nickname, Telephone: telephone, Password: password,
		RePassword: rePassword, Age: age}
	common.DB.Create(&user)
}


