package student

import (
	"context"
	"virtual-diary/internal/student/studentdao"
)

type StudentRepo interface {
	GetStudentById(id string, timeoutCtx context.Context) (studentdao.Student, error)
	GetStudentsByClass(className string, timeoutCtx context.Context) ([]studentdao.Student, error)
	CreateStudent(name string, surname string, age int, classId uint, timeoutCtx context.Context) (bool, error)
	DeleteStudent(studentId string, timeoutCtx context.Context) (bool, error)
}
