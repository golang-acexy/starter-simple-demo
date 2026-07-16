package consts

import (
	"time"

	"github.com/golang-acexy/starter-redis/redisstarter"
)

var (
	// RedisStudentKey 学生缓存10秒过期
	RedisStudentKey = redisstarter.NewRedisKey("student:%d", time.Second*10)
)
