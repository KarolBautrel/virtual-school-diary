package class

import (
	"fmt"

	"gorm.io/gorm"
)

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepoImpl{db: db}
}

func (s *classRepoImpl) GetAllClasses() string {
	return "repo for classes"
}

func (s *classRepoImpl) GetClassById(id string) (string, error) {

	return fmt.Sprintf("Shot for class with ID: %s", id), nil
}
