package student

import "virtual-diary/internal/student/studentdao"

type StudentRepo interface {
	GetStudentById(id string) (studentdao.Student, error)
	GetStudentsByClass(className string) ([]studentdao.Student, error)
	CreateStudent(name string, surname string, age int, classId uint) (bool, error)
	DeleteStudent(studentId string) (bool, error)
}
