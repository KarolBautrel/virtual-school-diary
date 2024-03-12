package class

import (
	"time"
	globalutils "virtual-diary/pkg/utils"
)

type ClassWriteService struct {
	repository ClassRepository
}

func NewWriteClassService(repo ClassRepository) *ClassWriteService {
	return &ClassWriteService{repository: repo}
}

func (s *ClassWriteService) CreateClass(name string, profile string) (bool, error) {
	status, err := s.repository.CreateClass(name, profile)
	if err != nil {
		return false, err
	}
	return status, err
}
func (s *ClassWriteService) RemoveClass(classId string) (bool, error) {
	ctx, close := globalutils.NewTimeoutContext(time.Microsecond * 10000)
	defer close()
	status, err := s.repository.RemoveClass(classId, ctx)
	if err != nil {
		return false, err
	}
	return status, err
}

func (s *ClassWriteService) RemoveStudentFromClass(studentId string, classId string) (bool, error) {
	status, err := s.repository.RemoveStudentFromClass(studentId, classId)
	if err != nil {
		return false, err
	}
	return status, err
}
