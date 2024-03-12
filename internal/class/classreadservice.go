package class

import (
	"time"
	"virtual-diary/internal/class/classdao"
	gloablutils "virtual-diary/pkg/utils"
)

type ClassReadService struct {
	repository ClassRepository
}

func NewReadClassService(repo ClassRepository) *ClassReadService {
	return &ClassReadService{repository: repo}
}

func (s *ClassReadService) GetAllClasses() ([]classdao.Class, error) {
	ctx, cancel := gloablutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()
	classes, err := s.repository.GetAllClasses(ctx)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *ClassReadService) GetClassById(id string) (classdao.Class, error) {
	ctx, cancel := gloablutils.NewTimeoutContext(time.Millisecond * 10000)
	defer cancel()

	class, err := s.repository.GetClassById(id, ctx)
	if err != nil {
		return classdao.Class{}, err
	}

	return class, nil
}
