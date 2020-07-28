package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shanshui/common"
	"shanshui/pkg/setting"
	"shanshui/pkg/utils"
	"shanshui/services/users"
)

func main() {
	db, err := common.InitDB()
	// 加载配置文件
	setting.Setup()

	if err != nil {
		fmt.Println("err open databases", err)
		return
	}
	common.CheckHasTable("users", &users.User{})
	defer db.Close()

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.StaticFS("/static", http.Dir("./static"))

	// eg: package.LoadPack(r)
	users.LoadUser(r)

	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}


