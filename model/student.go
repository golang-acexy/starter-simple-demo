package model

import "github.com/golang-acexy/starter-gorm/gormstarter"

const (
	SexMale Sex = iota
	SexFemale
)

type Sex uint8

type Student struct {
	ID         int64                 `gorm:"primaryKey;<-:false" json:"id"`
	CreateTime gormstarter.Timestamp `gorm:"<-:false" json:"createTime"`
	UpdateTime gormstarter.Timestamp `gorm:"<-:false" json:"updateTime"`
	Name       string                `json:"name"`
	Sex        Sex                   `json:"sex"`
	Age        uint8                 `json:"age"`
	TeacherID  int64                 `json:"teacherId"`
}

func (Student) TableName() string {
	return "demo_student"
}

type StudentMapper struct {
	gormstarter.BaseMapper[Student]
}
