package student

type studentWriteServiceImpl struct {
	repository StudentRepo
}

func NewWriteStudentService(repo StudentRepo) StudentWriteService {
	return &studentWriteServiceImpl{repository: repo}
}
