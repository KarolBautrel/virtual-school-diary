package student

type studentServiceImpl struct {
	repository StudentRepo
}

func NewReadStudentService(repo StudentRepo) StudentReadService {
	return &studentServiceImpl{repository: repo}
}
