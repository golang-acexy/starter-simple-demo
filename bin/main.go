package main

import (
	"github.com/acexy/golang-toolkit/logger"
	"github.com/acexy/golang-toolkit/sys"
	"github.com/acexy/starter-simple-demo/router"
	"github.com/golang-acexy/starter-gin/ginstarter"
	"github.com/golang-acexy/starter-gorm/gormstarter"
	"github.com/golang-acexy/starter-parent/parent"
	"github.com/golang-acexy/starter-redis/redisstarter"
	"github.com/redis/go-redis/v9"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func setLogger() {
	logger.EnableConsole(logger.TraceLevel, false)
	logger.EnableFileWithText(logger.TraceLevel, &lumberjack.Logger{
		Filename:   "./logs/log.log",
		MaxSize:    200,
		MaxBackups: 100,
		MaxAge:     365,
		Compress:   true,
	})
}

func loadStarter() []parent.Starter {
	return []parent.Starter{
		// gorm 组件
		&gormstarter.GormStarter{
			Config: gormstarter.GormConfig{
				Host:     "127.0.0.1",
				Port:     13306,
				Username: "root",
				Password: "root",
				Database: "test",
			},
		},
		// Redis 组件
		&redisstarter.RedisStarter{
			Config: redisstarter.RedisConfig{
				UniversalOptions: redis.UniversalOptions{
					Addrs:    []string{":6379"},
					Password: "tech-acexy",
					DB:       0,
				},
			},
		},

		// 加载ginStarter 最后启动 已保证http服务启动后 其他整体服务已可用
		&ginstarter.GinStarter{
			Config: ginstarter.GinConfig{
				ListenAddress: ":8080",
				DebugModule:   true,
				Routers: []ginstarter.Router{
					&router.StudentRouter{},
				},
			},
		},
	}
}

func main() {
	setLogger()

	starterLoader := parent.NewStarterLoader(loadStarter())
	err := starterLoader.Start()
	if err != nil {
		logger.Logrus().WithError(err).Errorln("应用启动异常")
		os.Exit(1)
	}

	sys.ShutdownHolding()
	_, _ = starterLoader.StopBySetting(time.Second * 30)
}
