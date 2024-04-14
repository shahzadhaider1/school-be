package service

import (
	"school-be/internal/repository/postgres"
	"school-be/pkg/models"
)

type ClassService struct {
	ClassRepository postgres.ClassRepository
}

func NewClassService() *ClassService {
	classRepo := postgres.NewClassRepository(postgres.DB)
	return &ClassService{
		ClassRepository: classRepo,
	}
}

type Class interface {
	CreateClass(class *models.Class) error
	GetClassByID(id int) (*models.Class, error)
	ListClasses(filter map[string]interface{}) ([]*models.Class, error)
	UpdateClass(class *models.Class) error
	DeleteClass(id int) error
}

func (s *ClassService) CreateClass(class *models.Class) error {
	return s.ClassRepository.CreateClass(class)
}

func (s *ClassService) GetClassByID(id int) (*models.Class, error) {
	return s.ClassRepository.GetClassByID(id)
}

func (s *ClassService) ListClasses(filter map[string]interface{}) ([]*models.Class, error) {
	return s.ClassRepository.ListClasses(filter)
}

func (s *ClassService) UpdateClass(class *models.Class) error {
	return s.ClassRepository.UpdateClass(class)
}

func (s *ClassService) DeleteClass(id int) error {
	return s.ClassRepository.DeleteClass(id)
}
