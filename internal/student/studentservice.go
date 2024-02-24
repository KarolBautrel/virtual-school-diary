package student

import (
	"gorm.io/gorm"
)

type studentServiceImpl struct {
	db         *gorm.DB
	repository StudentRepo
}

func NewStudentService(db *gorm.DB, repo StudentRepo) StudentService {
	return &studentServiceImpl{db: db, repository: repo}
}

// GetStudentById implements StudentService.
func (s *studentServiceImpl) GetStudentById(id string) (string, error) {
	student, err := s.repository.GetStudentById(id)
	if err == nil {
		panic("Student does not exists")
	}
	return student, err
}

// GetStudentsByClass implements StudentService.
func (s *studentServiceImpl) GetStudentsByClass(className string) (string, error) {
	students, err := s.repository.GetStudentsByClass(className)
	if err == nil {
		panic("Student does not exists")
	}
	return students, err
}
func (s *studentServiceImpl) GetWelcomeMessage() string {
	return "Welcome to the Student page!"
}
