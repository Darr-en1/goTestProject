package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var nilStruct = struct {
}{}

type Response struct {
	Data      interface{} `json:"data"`       // 返回结果  成功: 结果内容; 失败： {}
	ErrMsg    string      `json:"err_msg"`    // 错误信息
	IsSuccess bool        `json:"is_success"` // 请求结果 成功: true; 失败： false
}

const (
	ERROR   = false
	SUCCESS = true
)

func Result(httpCode int, data interface{}, errMsg string, isSuccess bool, c *gin.Context) {
	c.JSON(httpCode, Response{
		data,
		errMsg,
		isSuccess,
	})
}

func SuccessResult(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "", SUCCESS, c)
}

func BadRequestResult(errMsg string, c *gin.Context) {
	Result(http.StatusBadRequest, nilStruct, errMsg, ERROR, c)
}

func InternalServerErrorResult(errMsg string, c *gin.Context) {
	Result(http.StatusInternalServerError, nilStruct, errMsg, ERROR, c)
}

func Ok(c *gin.Context) {
	SuccessResult(map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	SuccessResult(data, c)
}

func Fail(c *gin.Context) {
	BadRequestResult("操作错误", c)
}

func FailWithMessage(errMsg string, c *gin.Context) {
	BadRequestResult(errMsg, c)
}

func InternalServerError(c *gin.Context) {
	InternalServerErrorResult("服务异常", c)
}

func InternalServerErrorWithMessage(errMsg string, c *gin.Context) {
	InternalServerErrorResult(errMsg, c)
}
