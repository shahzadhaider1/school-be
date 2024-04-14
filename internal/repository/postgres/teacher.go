package postgres

import (
	"fmt"

	"gorm.io/gorm"

	"school-be/pkg/models"
)

type TeacherRepository interface {
	CreateTeacher(teacher *models.Teacher) error
	GetTeacherByID(id int) (*models.Teacher, error)
	ListTeachers(filter map[string]interface{}) ([]*models.Teacher, error)
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(id int) error
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (repo *teacherRepository) CreateTeacher(teacher *models.Teacher) error {
	if err := repo.db.Create(teacher).Error; err != nil {
		fmt.Printf("failed to create teacher, error : %v", err)
		return err
	}

	return nil
}

func (repo *teacherRepository) GetTeacherByID(id int) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := repo.db.First(&teacher, id).Error; err != nil {
		fmt.Printf("failed to find teacher with ID : %v, error: %v", id, err)
		return nil, err
	}

	return &teacher, nil
}

func (repo *teacherRepository) ListTeachers(filter map[string]interface{}) ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	if err := repo.db.Where(filter).Find(&teachers).Error; err != nil {
		fmt.Printf("failed to fetch teachers, error : %v", err)

		return nil, err
	}

	return teachers, nil
}

func (repo *teacherRepository) UpdateTeacher(teacher *models.Teacher) error {
	result := repo.db.Where("id = ?", teacher.ID).Updates(teacher)

	if result.Error != nil {
		fmt.Printf("failed to update teacher with id : %v, error : %v", teacher.ID, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *teacherRepository) DeleteTeacher(id int) error {
	if err := repo.db.Where("id = ?", id).Delete(&models.Teacher{}).Error; err != nil {
		fmt.Printf("failed to delete teacher with id : %d, error : %v", id, err)
		return err
	}

	return nil
}
