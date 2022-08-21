package response

import (
	"RedBubble/common/responseCode"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code": 1000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}
*/
//通用响应对象

type ResponseData struct {
	Code responseCode.ResCode `json:"code"` // 程序中的错误码
	Msg  interface{}          `json:"msg"`  // 提示信息
	Data interface{}          `json:"data"` // 数据
}

func Error(c *gin.Context, code responseCode.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ErrorWithMsg(c *gin.Context, code responseCode.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: responseCode.CodeSuccess,
		Msg:  responseCode.CodeSuccess.Msg(),
		Data: data,
	})
}
