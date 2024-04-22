package class

import (
	"context"
	"virtual-diary/internal/class/classdao"
)

type ClassRepository interface {
	GetAllClasses(timeoutContext context.Context) ([]classdao.Class, error)
	GetClassById(id string, timeoutContext context.Context) (classdao.Class, error)
	CreateClass(name string, profile string) (bool, error)
	RemoveClass(classId string, timeoutContext context.Context) (bool, error)
	RemoveStudentFromClass(studentId string, classId string) (bool, error)
	AddStudentToClass(studentId string, classId string, timeoutContext context.Context) (bool, error)
}
