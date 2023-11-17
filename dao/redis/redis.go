package redis

import (
	"context"
	"fmt"
	"goWebCli/setting"

	"github.com/redis/go-redis/v9"
)

// 连接redis数据库

var rdb *redis.Client

// Init 初始化连接redis数据库
func Init(cfg *setting.RedisConfig) (err error) {
	// 从配置文件中获取连接信息，并连接数据库
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	// 测试连接是否成功
	_, err = rdb.Ping(context.Background()).Result()

	return
}

func Close() {
	_ = rdb.Close()
}
