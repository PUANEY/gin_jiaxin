package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@/shanshui?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		DB = db
	}
	return nil, err
}

func CheckHasTable(tablename string, data interface{})  {
	if !DB.HasTable(tablename) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(data)
	}
}