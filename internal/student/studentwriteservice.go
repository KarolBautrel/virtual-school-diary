package student

import (
	"fmt"
	"strconv"
	"time"
	globalutils "virtual-diary/pkg/utils"
)

type StudentWriteService struct {
	repository StudentRepo
}

func NewWriteStudentService(repo StudentRepo) *StudentWriteService {
	return &StudentWriteService{repository: repo}
}

func (s *StudentWriteService) CreateStudent(name, surname, age, classId string) (bool, error) {
	intAge, err := strconv.Atoi(age)
	ctx, cancel := globalutils.NewTimeoutContext(time.Microsecond * 1000)
	defer cancel()
	if err != nil {
		return false, fmt.Errorf("conversion to int failed: %w", err)
	}

	intId, err := strconv.Atoi(classId)
	if err != nil {
		return false, fmt.Errorf("conversion to int failed: %w", err)
	}

	return s.repository.CreateStudent(name, surname, intAge, uint(intId), ctx)
}

func (s *StudentWriteService) DeleteStudent(studentId string) (bool, error) {
	ctx, cancel := globalutils.NewTimeoutContext(time.Microsecond * 1000)
	defer cancel()

	_, err := s.repository.DeleteStudent(studentId, ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
