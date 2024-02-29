package class

type classReadServiceImpl struct {
	repository ClassRepository
}

func NewReadClassService(repo ClassRepository) ClassReadService {
	return &classReadServiceImpl{repository: repo}
}
