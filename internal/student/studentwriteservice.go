package student

import (
	"fmt"
	"strconv"
)

type studentWriteServiceImpl struct {
	repository StudentRepo
}

func NewWriteStudentService(repo StudentRepo) StudentWriteService {
	return &studentWriteServiceImpl{repository: repo}
}

func (s *studentWriteServiceImpl) CreateStudent(name, surname, age, classId string) (bool, error) {
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
