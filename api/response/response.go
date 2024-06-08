package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ResponseSuccess struct {
	Code int    `json:"code" example:"0"`
	Data string `json:"data" example:"{}"`
	Msg  string `json:"msg" example:"success"`
}

type ResponseFail struct {
	Code int    `json:"code" example:"-1"`
	Data string `json:"data" example:""`
	Msg  string `json:"msg" example:"fail reason"`
}

const (
	ERROR           = -1
	SUCCESS         = 0
	SUCCESS_MESSAGE = "success"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	httpStatus := http.StatusOK
	if code != 0 {
		httpStatus = http.StatusBadRequest
	}
	c.JSON(httpStatus, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "failed", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
