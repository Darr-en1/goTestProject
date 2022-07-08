package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data      interface{} `json:"data,omitempty"`     // 返回结果
	ErrMsg    string      `json:"err_msg,omitempty"`  // 错误信息
	ErrCode   errCode     `json:"err_code,omitempty"` // 错误状态码
	IsSuccess bool        `json:"is_success"`         // 请求结果
}

const (
	Fail    = false
	SUCCESS = true
)

func Result(httpCode int, data interface{}, errCode errCode, isSuccess bool, c *gin.Context) {
	c.JSON(httpCode, Response{
		data,
		errCode.String(),
		errCode,
		isSuccess,
	})
}

func ResultOk(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Data:      data,
		IsSuccess: SUCCESS,
	})
}

func ResultFail(httpCode int, errCode errCode, c *gin.Context) {
	c.JSON(httpCode, Response{
		ErrMsg:    errCode.String(),
		ErrCode:   errCode,
		IsSuccess: Fail,
	})
}
