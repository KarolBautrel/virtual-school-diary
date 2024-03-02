package class

import (
	"virtual-diary/internal/class/classdao"

	"gorm.io/gorm"
)

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepoImpl{db: db}
}

func (r *classRepoImpl) GetAllClasses() ([]classdao.Class, error) {
	var classes []classdao.Class
	err := r.db.Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepoImpl) GetClassById(id string) (classdao.Class, error) {
	class := classdao.Class{}
	result := r.db.Where("id = ?", id).First(&class)
	if result.Error != nil {
		return classdao.Class{}, result.Error
	}
	return class, nil
}

func (r *classRepoImpl) CreateClass(name string, profile string) (bool, error) {
	newClass := classdao.Class{
		Name:    name,
		Profile: profile,
	}
	result := r.db.Create(&newClass)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *classRepoImpl) RemoveClass(classId string) (bool, error) {
	class, err := r.GetClassById(classId)
	if err != nil {
		return false, err
	}

	if result := r.db.Delete(&class); result.Error != nil {
		return false, result.Error
	}

	return true, nil

}

func (r *classRepoImpl) RemoveStudentFromClass(studentId string, classId string) string {

	return "will be after student dao"
}
