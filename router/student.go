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
	router.GET("get/:id", s.getById())
}
func (*StudentRouter) save() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		var student model.Student
		request.MustBindBodyJson(&student)
		_, _ = studentMapper.Save(&student)
		return ginstarter.RespRestSuccess(student.ID), nil
	}
}

func (*StudentRouter) getById() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		//studentId := conversion.ParseIntPanic(request.GetPathParam("id"))
		studentId := conversion.ParseInt(request.GetPathParam("id"))
		var result model.Student
		err := redisStringCmd.GetAnyWithGob(consts.RedisStudentKey, &result, studentId)
		if err == nil {
			logger.Logrus().Debugln("已从缓存中获取学生信息", studentId)
		} else {
			row, err := studentMapper.SelectById(studentId, &result)
			if err != nil {
				return nil, err
			} else if row == 0 {
				return ginstarter.RespRestBadParameters("无效的学生ID"), nil
			} else {
				err = redisStringCmd.SetAnyWithGob(consts.RedisStudentKey, result, studentId)
				if err != nil {
					logger.Logrus().Errorln("设置学生信息缓存失败", studentId)
					return nil, err
				} else {
					logger.Logrus().Debugln("已从数据库中获取学生信息并创建缓存", studentId)
				}
			}
		}
		return ginstarter.RespRestSuccess(result), nil
	}
}
