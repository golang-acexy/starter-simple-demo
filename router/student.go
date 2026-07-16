package router

import (
	"github.com/acexy/golang-toolkit/logger"
	"github.com/acexy/golang-toolkit/math/conversion"
	"github.com/acexy/starter-simple-demo/consts"
	"github.com/acexy/starter-simple-demo/model"
	"github.com/golang-acexy/starter-gin/ginstarter"
	"github.com/golang-acexy/starter-redis/redisstarter"
)

var studentMapper model.StudentMapper
var redisStringCmd = redisstarter.StringCmd()

type StudentRouter struct {
}

func (*StudentRouter) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "student",
	}
}

func (s *StudentRouter) Handlers(router *ginstarter.RouterWrapper) {
	router.POST("save", s.save())
	router.GET("get/:id", s.getByID())
}
func (*StudentRouter) save() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		var student model.Student
		request.MustBindBodyJSON(&student)
		if _, err := studentMapper.Insert(&student); err != nil {
			return nil, err
		}
		return ginstarter.RespRestSuccess(student.ID), nil
	}
}

func (*StudentRouter) getByID() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		studentID := conversion.ParseInt(request.GetPathParam("id"))
		if studentID <= 0 {
			return ginstarter.RespRestBadParameters("无效的学生ID"), nil
		}
		var result model.Student
		err := redisStringCmd.GetAnyWithGob(consts.RedisStudentKey, &result, studentID)
		if err == nil {
			logger.Logrus().Debugln("已从缓存中获取学生信息", studentID)
		} else {
			logger.Logrus().WithError(err).Debugln("学生缓存未命中，回源数据库", studentID)
			row, err := studentMapper.SelectByID(studentID, &result)
			if err != nil {
				return nil, err
			} else if row == 0 {
				return ginstarter.RespRestBadParameters("无效的学生ID"), nil
			} else {
				err = redisStringCmd.SetAnyWithGob(consts.RedisStudentKey, result, studentID)
				if err != nil {
					logger.Logrus().WithError(err).Errorln("设置学生信息缓存失败", studentID)
				} else {
					logger.Logrus().Debugln("已从数据库中获取学生信息并创建缓存", studentID)
				}
			}
		}
		return ginstarter.RespRestSuccess(result), nil
	}
}
