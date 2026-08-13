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

// StudentColumns 定义 Student Mapper 可使用的类型安全字段。
type StudentColumns struct {
	ID        gormstarter.Column[Student, int64]
	Name      gormstarter.Column[Student, string]
	Sex       gormstarter.Column[Student, Sex]
	Age       gormstarter.Column[Student, uint8]
	TeacherID gormstarter.Column[Student, int64]
}

var studentColumns = StudentColumns{
	ID:        gormstarter.NewColumn(func(student *Student) *int64 { return &student.ID }),
	Name:      gormstarter.NewColumn(func(student *Student) *string { return &student.Name }),
	Sex:       gormstarter.NewColumn(func(student *Student) *Sex { return &student.Sex }),
	Age:       gormstarter.NewColumn(func(student *Student) *uint8 { return &student.Age }),
	TeacherID: gormstarter.NewColumn(func(student *Student) *int64 { return &student.TeacherID }),
}

// Columns 返回 Student Mapper 的类型安全字段集合。
func (StudentMapper) Columns() *StudentColumns {
	return &studentColumns
}
