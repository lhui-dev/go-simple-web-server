// 通用返回结果
// author: 李少辉

package result

import "github.com/gin-gonic/gin"

// Result 通用返回结构
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回的数据
}

// Success 返回成功消息体
func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	ctx.JSON(res.Code, res)
}

func SuccessWithCodeAndMessage(ctx *gin.Context, code uint, message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(code)
	res.Message = message
	res.Data = data
	ctx.JSON(res.Code, res)
}

func SuccessWithMessage(ctx *gin.Context, message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = message
	res.Data = data
	ctx.JSON(res.Code, res)
}

// Failed 返回失败消息体
func Failed(ctx *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	ctx.JSON(code, res)
}
