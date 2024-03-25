package student

import (
	"virtual-diary/internal/student/studentdao"
	"virtual-diary/internal/student/studentdto"
)

func ConvertDaoToDto(studentDAO *studentdao.Student, studentDTO *studentdto.StudentDTO) {
	studentDTO.Name = studentDAO.Name
	studentDTO.Surname = studentDAO.Surname
	studentDTO.Age = studentDAO.Age
}
