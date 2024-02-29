package class

type classRepoImpl struct {
}

func NewClassRepository() ClassRepository {
	return &classRepoImpl{}
}
