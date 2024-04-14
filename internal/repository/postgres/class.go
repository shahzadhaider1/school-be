package postgres

import (
	"fmt"

	"gorm.io/gorm"

	"school-be/pkg/models"
)

type ClassRepository interface {
	CreateClass(class *models.Class) error
	GetClassByID(id int) (*models.Class, error)
	ListClasses(filter map[string]interface{}) ([]*models.Class, error)
	UpdateClass(class *models.Class) error
	DeleteClass(id int) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (repo *classRepository) CreateClass(class *models.Class) error {
	if err := repo.db.Create(class).Error; err != nil {
		fmt.Printf("failed to create class, error : %v", err)
		return err
	}

	return nil
}

func (repo *classRepository) GetClassByID(id int) (*models.Class, error) {
	var class models.Class
	if err := repo.db.First(&class, id).Error; err != nil {
		fmt.Printf("failed to find class with ID : %v, error: %v", id, err)
		return nil, err
	}

	return &class, nil
}

func (repo *classRepository) ListClasses(filter map[string]interface{}) ([]*models.Class, error) {
	var classs []*models.Class
	if err := repo.db.Where(filter).Find(&classs).Error; err != nil {
		fmt.Printf("failed to fetch classs, error : %v", err)

		return nil, err
	}

	return classs, nil
}

func (repo *classRepository) UpdateClass(class *models.Class) error {
	result := repo.db.Where("id = ?", class.ID).Updates(class)

	if result.Error != nil {
		fmt.Printf("failed to update class with id : %v, error : %v", class.ID, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *classRepository) DeleteClass(id int) error {
	if err := repo.db.Where("id = ?", id).Delete(&models.Class{}).Error; err != nil {
		fmt.Printf("failed to delete class with id : %d, error : %v", id, err)
		return err
	}

	return nil
}
