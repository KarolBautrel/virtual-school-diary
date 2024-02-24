package class

import (
	"fmt"

	"gorm.io/gorm"
)

type classServiceImpl struct {
	db         *gorm.DB
	repository ClassRepository
}

func NewClassService(db *gorm.DB, repo ClassRepository) ClassService {
	return &classServiceImpl{db: db, repository: repo}
}

func (s *classServiceImpl) GetAllClasses() string {
	return s.repository.GetAllClasses()

}

func (s *classServiceImpl) GetClassById(id string) (string, error) {
	class, err := s.repository.GetClassById(id)
	if err == nil {
		panic("ELO")
	}
	fmt.Println(class)
	return fmt.Sprintf("Shot for class with ID: %s", id), nil
}

func (s *classServiceImpl) GetWelcomeMessage() string {
	return "Welcome to the Class page!"
}
