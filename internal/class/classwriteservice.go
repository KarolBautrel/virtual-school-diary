package class

type classWriteServiceImpl struct {
	repository ClassRepository
}

func NewWriteClassService(repo ClassRepository) ClassWriteService {
	return &classWriteServiceImpl{repository: repo}
}

func (s *classWriteServiceImpl) CreateClass(name string, profile string) (bool, error) {
	status, err := s.repository.CreateClass(name, profile)
	if err != nil {
		return false, err
	}
	return status, err
}
func (s *classWriteServiceImpl) RemoveClass(classId string) (bool, error) {
	status, err := s.repository.RemoveClass(classId)
	if err != nil {
		return false, err
	}
	return status, err
}

func (s *classWriteServiceImpl) RemoveStudentFromClass(studentId string, classId string) (bool, error) {
	status, err := s.repository.RemoveStudentFromClass(studentId, classId)
	if err != nil {
		return false, err
	}
	return status, err
}
