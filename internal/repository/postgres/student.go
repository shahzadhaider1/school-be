package postgres

import (
	"fmt"

	"gorm.io/gorm"

	"school-be/pkg/models"
)

type StudentRepository interface {
	CreateStudent(Student *models.Student) error
	GetStudentByID(id int) (*models.Student, error)
	ListStudents(filter map[string]interface{}) ([]*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id int) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (repo *studentRepository) CreateStudent(student *models.Student) error {
	if err := repo.db.Create(student).Error; err != nil {
		fmt.Printf("failed to create student, error : %v", err)
		return err
	}

	return nil
}

func (repo *studentRepository) GetStudentByID(id int) (*models.Student, error) {
	var student models.Student
	if err := repo.db.First(&student, id).Error; err != nil {
		fmt.Printf("failed to find student with ID : %v, error: %v", id, err)
		return nil, err
	}

	return &student, nil
}

func (repo *studentRepository) ListStudents(filter map[string]interface{}) ([]*models.Student, error) {
	var students []*models.Student
	if err := repo.db.Where(filter).Find(&students).Error; err != nil {
		fmt.Printf("failed to fetch students, error : %v", err)

		return nil, err
	}

	return students, nil
}

func (repo *studentRepository) UpdateStudent(student *models.Student) error {
	result := repo.db.Where("id = ?", student.ID).Updates(student)

	if result.Error != nil {
		fmt.Printf("failed to update student with id : %v, error : %v", student.ID, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *studentRepository) DeleteStudent(id int) error {
	if err := repo.db.Where("id = ?", id).Delete(&models.Student{}).Error; err != nil {
		fmt.Printf("failed to delete student with id : %d, error : %v", id, err)
		return err
	}

	return nil
}
