package classdao

import (
	"virtual-diary/internal/student/studentdao"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name     string
	Profile  string
	Students []studentdao.Student
}
