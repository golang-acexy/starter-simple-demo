package model

import "github.com/golang-acexy/starter-gorm/gormstarter"

const (
	SexMale Sex = iota
	SexFemale
)

type Sex uint8

type Student struct {
	gormstarter.BaseModel[int]
	CreateTime gormstarter.Timestamp `gorm:"<-:false" json:"createTime"`
	UpdateTime gormstarter.Timestamp `gorm:"<-:update" json:"updateTime"`
	Name       string                `json:"name"`
	Sex        Sex                   `json:"sex"`
	Age        uint8                 `json:"age"`
	TeacherId  int                   `json:"teacherId"`
}

func (Student) TableName() string {
	return "demo_student"
}

type StudentMapper struct {
	gormstarter.BaseMapper[Student]
}
