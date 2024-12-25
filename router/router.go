// gin路由信息
// author 李少辉

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-simple-web-server/internal/middleware"
	"net/http"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("color.html")

	// 宕机恢复
	router.Use(gin.Recovery())
	// 日志中间件
	router.Use(middleware.Logger())
	// 跨域中间件
	router.Use(middleware.Cors())

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "color.html", gin.H{
			"color": "red",
		})
	})

	return router
}
