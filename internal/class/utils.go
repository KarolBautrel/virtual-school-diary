package class

import (
	"virtual-diary/internal/class/classdao"
	"virtual-diary/internal/class/classdto"
)

func ConvertClassDaoToDto(classDTO *classdto.ClassDto, classDAO *classdao.Class) {
	classDTO.Name = classDAO.Name
	classDTO.Profile = classDAO.Profile
}
