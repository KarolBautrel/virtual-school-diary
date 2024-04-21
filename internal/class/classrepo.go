package class

import (
	"context"
	"fmt"
	"virtual-diary/internal/class/classdao"

	"gorm.io/gorm"
)

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepoImpl{db: db}
}

func (r *classRepoImpl) GetAllClasses(timeoutContext context.Context) ([]classdao.Class, error) {
	var classes []classdao.Class
	err := r.db.WithContext(timeoutContext).Find(&classes).Error
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepoImpl) GetClassById(id string, timeoutContext context.Context) (classdao.Class, error) {
	class := classdao.Class{}
	result := r.db.WithContext(timeoutContext).Where("id = ?", id).First(&class)
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

func (r *classRepoImpl) RemoveClass(classId string, timeoutContext context.Context) (bool, error) {
	class, err := r.GetClassById(classId, timeoutContext)
	if err != nil {
		return false, err
	}

	if result := r.db.WithContext(timeoutContext).Delete(&class); result.Error != nil {
		return false, result.Error
	}

	return true, nil

}

func (r *classRepoImpl) RemoveStudentFromClass(studentId string, classId string, timeoutContext context.Context) (bool, error) {
	fmt.Println("After Creation of student domain there will be something here")
	return true, nil
}

func (r *classRepoImpl) AddStudentToClass(studentId string, classId string, timeoutContext context.Context) (bool, error) {
	///Impl coming soon
	return true, nil
}
