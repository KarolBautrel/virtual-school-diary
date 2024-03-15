package studentdao

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	Age     int    `gorm:"type:int;not null"`
	Surname string `gorm:"type:varchar(100);not null"`
	ClassID uint   `gorm:"type:int;not null"`
}
