package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/leedev/go-simple-web-server/docs"
	"github.com/leedev/go-simple-web-server/internal/configuration"
	"github.com/leedev/go-simple-web-server/internal/pkg/log"
	"github.com/leedev/go-simple-web-server/internal/pkg/redis"
	"github.com/leedev/go-simple-web-server/router"
)

// 加载日志log
var _log = log.Log()

// @title web服务接口文档
// @version 1.0
// @description web服务API接口文档

// @contact.name	Lee
// @contact.url	http://www.swagger.io/support
// @contact.email	lee.dev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9999
// @BasePath  /api/v1
func main() {

	gin.SetMode(configuration.Config.Server.Mode)
	// fmt.Println("configuration information: ", configuration.Config)

	// 初始化路由
	_router := router.InitRouter()

	// 服务器地址 host:port
	addr := fmt.Sprintf("%s:%s", configuration.Config.Server.Host, configuration.Config.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: _router,
	}

	// @see detail at : https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			_log.Fatalf("Listen at: %s", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_log.Info("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		_log.Fatalf("Shutdown server: %v\n", err)
	}
	select {
	case <-ctx.Done():
		_log.Println("Timeout of 3 seconds.")
	}
	_log.Println("Server exiting...")
}

// 初始化连接信息
func init() {
	// init mysql
	//if err := mysql_db.SetupMysqlDbLink(); err != nil {
	//	_log.Error(err)
	//}

	// init redis
	if err := redis.SetupRedisLink(); err != nil {
		_log.Error(err)
	}
}
