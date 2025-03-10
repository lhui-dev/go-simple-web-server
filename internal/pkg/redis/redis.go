// Redis连接配置
// author: 李少辉

package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/leedev/go-simple-web-server/internal/configuration"
	"github.com/leedev/go-simple-web-server/internal/pkg/log"
)

var RedisClient *redis.Client

// SetupRedisLink 初始化redis连接
func SetupRedisLink() error {
	var logger = log.Log()
	var ctx = context.Background()

	var redisAddr = fmt.Sprintf("%v:%v", configuration.Config.Db.RedisConfig.Host, configuration.Config.Db.RedisConfig.Port)
	var redisPasswd = configuration.Config.Db.RedisConfig.Password
	var defaultRedisDb = configuration.Config.Db.RedisConfig.Database
	logger.Infof("redisAddr=[%v] redisPasswd=[%v] defaultRedisDb=[%v]", redisAddr, redisPasswd, defaultRedisDb)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPasswd,
		DB:       defaultRedisDb,
	})
	logger.Infof("redisDb=[%v]", RedisClient)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		return err
	}
	return nil
}
