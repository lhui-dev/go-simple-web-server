// gin路由信息
// author 李少辉

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-simple-web-server/api/controller"
	"github.com/leedev/go-simple-web-server/api/dao"
	"github.com/leedev/go-simple-web-server/api/service"
	"github.com/leedev/go-simple-web-server/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	router := gin.New()

	// 宕机恢复
	router.Use(gin.Recovery())
	// 日志中间件
	router.Use(middleware.Logger())
	// 跨域中间件
	router.Use(middleware.Cors())

	// dao
	userDao := dao.NewUserDaoImpl()

	// service
	userService := service.NewUserServiceImpl(userDao)

	// controller
	userController := controller.NewUserController(userService)

	// 注册路由
	registerRoute(router, userController)
	// 返回路由信息
	return router
}

// 注册路由信息
func registerRoute(
	router *gin.Engine,
	userController *controller.UserController,
) {

	// swagger doc router
	// --> /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRoutes := router.Group("/api/v1")
	{
		// 用户相关路由信息
		userRoutes := baseRoutes.Group("/user")
		{
			userRoutes.GET("/list", userController.QueryUserList)
		}
	}
}
