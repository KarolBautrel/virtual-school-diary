package studentdao

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string
	Age     int
	Surname string
	ClassID uint
}
