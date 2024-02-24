package student

type StudentService interface {
	GetWelcomeMessage() string
	GetStudentById(id string) (string, error)
	GetStudentsByClass(className string) (string, error)
}

type StudentRepo interface {
	GetStudentById(id string) (string, error)
	GetStudentsByClass(className string) (string, error)
}
