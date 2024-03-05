package class

import "virtual-diary/internal/class/classdao"

type classReadServiceImpl struct {
	repository ClassRepository
}

func NewReadClassService(repo ClassRepository) ClassReadService {
	return &classReadServiceImpl{repository: repo}
}

func (s *classReadServiceImpl) GetAllClasses() ([]classdao.Class, error) {
	classes, err := s.repository.GetAllClasses()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *classReadServiceImpl) GetClassById(id string) (classdao.Class, error) {
	class, err := s.repository.GetClassById(id)
	if err != nil {
		return classdao.Class{}, err
	}

	return class, nil
}

func (s *classReadServiceImpl) GetWelcomeMessage() string {
	return "Welcome to the Class page!"
}
