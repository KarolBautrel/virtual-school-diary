package student

import (
	"fmt"
	"strconv"
)

type StudentWriteService struct {
	repository StudentRepo
}

func NewWriteStudentService(repo StudentRepo) *StudentWriteService {
	return &StudentWriteService{repository: repo}
}

func (s *StudentWriteService) CreateStudent(name, surname, age, classId string) (bool, error) {
	intAge, err := strconv.Atoi(age)
	if err != nil {
		return false, fmt.Errorf("conversion to int failed: %w", err)
	}

	intId, err := strconv.Atoi(classId)
	if err != nil {
		return false, fmt.Errorf("conversion to int failed: %w", err)
	}

	return s.repository.CreateStudent(name, surname, intAge, uint(intId))
}

func (s *StudentWriteService) DeleteStudent(studentId string) (bool, error) {
	_, err := s.repository.DeleteStudent(studentId)
	if err != nil {
		return false, err
	}
	return true, nil
}
