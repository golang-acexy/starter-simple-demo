package consts

import (
	"github.com/golang-acexy/starter-redis/redisstarter"
	"time"
)

var (
	// RedisStudentKey 学生缓存10秒过期
	RedisStudentKey = redisstarter.NewRedisKey("student:%d", time.Second*10)
)
