package student

import (
	"context"
	"virtual-diary/internal/student/studentdao"

	"gorm.io/gorm"
)

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepoImpl{db: db}
}

func (r *studentRepoImpl) GetStudentById(id string, timeoutCtx context.Context) (studentdao.Student, error) {
	student := studentdao.Student{}
	result := r.db.WithContext(timeoutCtx).Where("id = ?", id).First(&student)
	if result.Error != nil {
		return studentdao.Student{}, result.Error
	}

	return student, nil
}
func (r *studentRepoImpl) GetStudentsByClass(classId string, timeoutCtx context.Context) ([]studentdao.Student, error) {
	students := []studentdao.Student{}
	err := r.db.WithContext(timeoutCtx).Find(&students).Error
	if err != nil {
		return []studentdao.Student{}, err
	}
	return students, nil
}

func (r *studentRepoImpl) CreateStudent(name string, surname string, age int, classId uint, timeoutCtx context.Context) (bool, error) {

	newStudent := studentdao.Student{
		Name:    name,
		Surname: surname,
		Age:     age,
		ClassID: classId,
	}
	if result := r.db.WithContext(timeoutCtx).Create(&newStudent); result.Error != nil {
		return false, result.Error
	}
	return true, nil

}

func (r *studentRepoImpl) DeleteStudent(studentId string, timeoutCtx context.Context) (bool, error) {
	student, err := r.GetStudentById(studentId, timeoutCtx)
	if err != nil {
		return false, err
	}
	if result := r.db.WithContext(timeoutCtx).Delete(&student); result.Error != nil {
		return false, result.Error
	}
	return true, nil

}
