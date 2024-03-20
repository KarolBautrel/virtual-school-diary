package student

import (
	"virtual-diary/internal/student/studentdao"

	"gorm.io/gorm"
)

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return &studentRepoImpl{db: db}
}

func (r *studentRepoImpl) GetStudentById(id string) (studentdao.Student, error) {
	student := studentdao.Student{}
	result := r.db.Where("id = ?", id).First(&student)
	if result.Error != nil {
		return studentdao.Student{}, result.Error
	}

	return student, nil
}
func (r *studentRepoImpl) GetStudentsByClass(classId string) ([]studentdao.Student, error) {
	students := []studentdao.Student{}
	err := r.db.Find(&students).Error
	if err != nil {
		return []studentdao.Student{}, err
	}
	return students, nil
}

func (r *studentRepoImpl) CreateStudent(name string, surname string, age int, classId uint) (bool, error) {

	newStudent := studentdao.Student{
		Name:    name,
		Surname: surname,
		Age:     age,
		ClassID: classId,
	}
	if result := r.db.Create(&newStudent); result.Error != nil {
		return false, result.Error
	}
	return true, nil

}

func (r *studentRepoImpl) DeleteStudent(studentId string) (bool, error) {
	student, err := r.GetStudentById(studentId)
	if err != nil {
		return false, err
	}
	if result := r.db.Delete(&student); result.Error != nil {
		return false, result.Error
	}
	return true, nil

}
