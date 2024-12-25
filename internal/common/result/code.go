// 返回状态码
// author: 李少辉

package result

import "net/http"

// 自定义返回状态信息
type statusCode struct {
	SUCCESS      int
	FAILED       int
	NoAuth       int
	Unauthorized int
	BadRequest   int

	Message map[int]string
}

// ApiCode api 状态码
var ApiCode = &statusCode{
	SUCCESS:      http.StatusOK,
	FAILED:       http.StatusInternalServerError,
	NoAuth:       http.StatusForbidden,
	Unauthorized: http.StatusUnauthorized,
	BadRequest:   http.StatusBadRequest,
}

// 初始化状态信息
func init() {
	ApiCode.Message = map[int]string{
		ApiCode.SUCCESS:      "请求成功",
		ApiCode.FAILED:       "请求失败",
		ApiCode.NoAuth:       "权限不足",
		ApiCode.Unauthorized: "未认证",
		ApiCode.BadRequest:   "Bad Request",
	}
}

// GetMessage 供外部调用
func (c *statusCode) GetMessage(code int) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
