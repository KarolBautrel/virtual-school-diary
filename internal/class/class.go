package class

type ClassService interface {
	GetWelcomeMessage() string
	GetAllClasses() string
	GetClassById(id string) (string, error)
}

type ClassRepository interface {
	GetAllClasses() string
	GetClassById(id string) (string, error)
}
