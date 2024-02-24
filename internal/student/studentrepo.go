package student

import (
	"fmt"

	"gorm.io/gorm"
)

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepoImpl{db: db}
}

func (s *studentRepoImpl) GetStudentById(id string) (string, error) {
	return fmt.Sprintf("Shot for student with ID: %s", id), nil
}
func (s *studentRepoImpl) GetStudentsByClass(className string) (string, error) {
	panic("unimplemented")
}
