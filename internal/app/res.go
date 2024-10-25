package app

import (
	"cleangin/internal/e"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type FailResponse struct {
	Err string `json:"error"`
}

type ResponseMsg struct {
	Msg string `json:"msg"`
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) FailResponse(httpCode int, errCode int) {
	g.C.JSON(httpCode, FailResponse{
		Err: e.GetMsg(errCode),
	})
	return
}

func (g *Gin) Response(httpCode, msgCode int, data interface{}) {
	if data == nil {
		g.C.JSON(httpCode, ResponseMsg{
			Msg: e.GetMsg(msgCode),
		})
	} else {
		g.C.JSON(httpCode, Response{
			Msg:  e.GetMsg(msgCode),
			Data: data,
		})
	}
	return
}

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.ERROR
	}

	return http.StatusOK, e.SUCCESS
}

func GetUintParam(c *gin.Context, param string) (uint, error) {
	propString := c.Param(param)

	if propString == "" {
		return 0, fmt.Errorf("prop is not defined")
	}

	p, err := strconv.Atoi(propString)
	if err != nil {
		return 0, err
	}
	return uint(p), nil
}
