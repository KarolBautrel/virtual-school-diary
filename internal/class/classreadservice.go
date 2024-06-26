package class

import (
	"fmt"
	"time"
	"virtual-diary/internal/class/classdto"
	gloablutils "virtual-diary/pkg/utils"
)

type ClassReadService struct {
	repository ClassRepository
}

func NewReadClassService(repo ClassRepository) *ClassReadService {
	return &ClassReadService{repository: repo}
}

func (s *ClassReadService) GetAllClasses() ([]classdto.ClassDto, error) {
	ctx, cancel := gloablutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	classes, err := s.repository.GetAllClasses(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in read service: %s", err)
	}
	classesDTO := []classdto.ClassDto{}

	for _, classDAO := range classes {
		var classDTO classdto.ClassDto
		ConvertClassDaoToDto(&classDTO, &classDAO)
		classesDTO = append(classesDTO, classDTO)
	}
	return classesDTO, nil
}

func (s *ClassReadService) GetClassById(id string) (classdto.ClassDto, error) {
	ctx, cancel := gloablutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()

	classDAO, err := s.repository.GetClassById(id, ctx)
	var classDTO classdto.ClassDto
	if err != nil {
		return classdto.ClassDto{}, fmt.Errorf("error in read service: %s", err)

	}
	ConvertClassDaoToDto(&classDTO, &classDAO)

	return classDTO, nil
}
