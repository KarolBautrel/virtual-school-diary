package classdao

import (
	"virtual-diary/internal/student/studentdao"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name     string               `gorm:"type:varchar(100);not null"`
	Profile  string               `gorm:"type:varchar(100);not null"`
	Students []studentdao.Student `gorm:"foreignKey:ClassID"`
}
