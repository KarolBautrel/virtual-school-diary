package student

import (
	"gorm.io/gorm"
)

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepoImpl{db: db}
}
