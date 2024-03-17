package student

import (
	"virtual-diary/internal/student/studentdao"
)

type studentServiceImpl struct {
	repository StudentRepo
}

func NewReadStudentService(repo StudentRepo) StudentReadService {
	return &studentServiceImpl{repository: repo}
}

func (s *studentServiceImpl) GetStudentById(id string) (studentdao.Student, error) {
	student, err := s.repository.GetStudentById(id)
	if err != nil {
		return student, err
	}
	return student, err
}

func (s *studentServiceImpl) GetStudentsByClass(className string) ([]studentdao.Student, error) {
	students, err := s.repository.GetStudentsByClass(className)
	if err != nil {
		return students, err
	}
	return students, err
}
func (s *studentServiceImpl) GetWelcomeMessage() string {
	return "Welcome to the Student page!"
}
