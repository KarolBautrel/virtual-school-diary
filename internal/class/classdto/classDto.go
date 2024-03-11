package classdto

import "virtual-diary/internal/student/studentdto"

type ClassDto struct {
	Name     string
	Profile  string
	Students []studentdto.StudentDTO
}
