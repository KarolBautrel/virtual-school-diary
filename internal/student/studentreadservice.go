package student

import (
	"virtual-diary/internal/student/studentdao"
)

type StudentReadService struct {
	repository StudentRepo
}

func NewReadStudentService(repo StudentRepo) *StudentReadService {
	return &StudentReadService{repository: repo}
}

func (s *StudentReadService) GetStudentById(id string) (studentdao.Student, error) {
	student, err := s.repository.GetStudentById(id)
	if err != nil {
		return student, err
	}
	return student, err
}

func (s *StudentReadService) GetStudentsByClass(className string) ([]studentdao.Student, error) {
	students, err := s.repository.GetStudentsByClass(className)
	if err != nil {
		return students, err
	}
	return students, err
}
func (s *StudentReadService) GetWelcomeMessage() string {
	return "Welcome to the Student page!"
}
