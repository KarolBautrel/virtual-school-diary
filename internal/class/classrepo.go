package class

import (
	"context"
	"virtual-diary/internal/class/classdao"
	"virtual-diary/internal/student/studentdao"

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

func (r *classRepoImpl) RemoveStudentFromClass(studentId string, classId string) (bool, error) {
	var student studentdao.Student
	if err := r.db.Where("id = ? AND class_id = ?", studentId, classId).First(&student).Error; err != nil {
		return false, err
	}
	student.ClassID = 0
	if err := r.db.Save(&student).Error; err != nil {
		return false, err
	}

	return true, nil
}
func (r *classRepoImpl) AddStudentToClass(studentId string, classId string, timeoutContext context.Context) (bool, error) {
	var student studentdao.Student
	if err := r.db.Where("id = ?", studentId).First(&student).Error; err != nil {
		return false, err
	}

	class, err := r.GetClassById(classId, timeoutContext)
	if err != nil {
		return false, err
	}

	class.Students = append(class.Students, student)
	if err := r.db.Save(&class).Error; err != nil {
		return false, err
	}
	return true, nil
}
