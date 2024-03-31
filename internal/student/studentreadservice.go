package student

import (
	"virtual-diary/internal/student/studentdto"
)

type StudentReadService struct {
	repository StudentRepo
}

func NewReadStudentService(repo StudentRepo) *StudentReadService {
	return &StudentReadService{repository: repo}
}

func (s *StudentReadService) GetStudentById(id string) (studentdto.StudentDTO, error) {
	student, err := s.repository.GetStudentById(id)
	if err != nil {
		return studentdto.StudentDTO{}, err
	}
	var studentDTO studentdto.StudentDTO
	ConvertStudentDaoToDto(&student, &studentDTO)
	return studentDTO, err
}

func (s *StudentReadService) GetStudentsByClass(className string) ([]studentdto.StudentDTO, error) {
	students, err := s.repository.GetStudentsByClass(className)
	if err != nil {
		return []studentdto.StudentDTO{}, err
	}
	studentsDTO := []studentdto.StudentDTO{}

	for _, student := range students {
		var studentDTO studentdto.StudentDTO
		ConvertStudentDaoToDto(&student, &studentDTO)
		studentsDTO = append(studentsDTO, studentDTO)
	}
	return studentsDTO, err
}
func (s *StudentReadService) GetWelcomeMessage() string {
	return "Welcome to the Student page!"
}
