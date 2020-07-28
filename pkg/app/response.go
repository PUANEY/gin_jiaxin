package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"errcode"`
	Msg  string      `json:"errmsg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
	return
}

func (g *Gin) Error(v string) {
	g.C.JSON(http.StatusOK, Response{
		Code: http.StatusBadRequest,
		Msg:  v,
	})
	return
}

