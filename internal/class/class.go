package class

import "virtual-diary/internal/class/classdao"

type ClassReadService interface {
	GetAllClasses() ([]classdao.Class, error)
	GetClassById(id string) (classdao.Class, error)
}

type ClassWriteService interface {
}

type ClassRepository interface {
	GetAllClasses() ([]classdao.Class, error)
	GetClassById(id string) (classdao.Class, error)
	CreateClass(name string, profile string) (bool, error)
	RemoveClass(classId string) (bool, error)
	RemoveStudentFromClass(studentId string, classId string) (bool, error)
}
