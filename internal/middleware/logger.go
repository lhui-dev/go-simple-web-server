// gin日志中间件
// author: 李少辉

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-simple-web-server/internal/pkg/log"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

// Logger 日志处理
func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(ctx *gin.Context) {
		startTime := time.Now()

		// 设置appId变量
		ctx.Set("appId", "go-simple-web-server")

		// 放行
		ctx.Next()
		latencyTime := time.Since(startTime)

		// 请求方法
		requestMethod := ctx.Request.Method
		requestURI := ctx.Request.RequestURI
		requestHeader := ctx.Request.Header
		proto := ctx.Request.Proto
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		err := ctx.Err()
		body, _ := io.ReadAll(ctx.Request.Body)

		logger.WithFields(logrus.Fields{
			"response_status_code": statusCode,
			"latency_time":         latencyTime,
			"client_ip":            clientIP,
			"request_method":       requestMethod,
			"request_uri":          requestURI,
			"header":               requestHeader,
			"proto":                proto,
			"err":                  err,
			"body":                 body,
		}).Info()

	}
}
