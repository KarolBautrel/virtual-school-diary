package student

import (
	"time"
	"virtual-diary/internal/student/studentdto"
	globalutils "virtual-diary/pkg/utils"
)

type StudentReadService struct {
	repository StudentRepo
}

func NewReadStudentService(repo StudentRepo) *StudentReadService {
	return &StudentReadService{repository: repo}
}

func (s *StudentReadService) GetStudentById(id string) (studentdto.StudentDTO, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 1000)
	defer cancel()
	student, err := s.repository.GetStudentById(id, ctx)
	if err != nil {
		return studentdto.StudentDTO{}, err
	}
	var studentDTO studentdto.StudentDTO
	ConvertStudentDaoToDto(&student, &studentDTO)
	return studentDTO, err
}

func (s *StudentReadService) GetStudentsByClass(className string) ([]studentdto.StudentDTO, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Millisecond * 1000)
	defer cancel()
	students, err := s.repository.GetStudentsByClass(className, ctx)
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
