package class

type classWriteServiceImpl struct {
	repository ClassRepository
}

func NewWriteClassService(repo ClassRepository) ClassWriteService {
	return &classWriteServiceImpl{repository: repo}
}
