package service

import (
	"school-be/internal/repository/postgres"
	"school-be/pkg/models"
)

type TeacherService struct {
	TeacherRepository postgres.TeacherRepository
}

func NewTeacherService() *TeacherService {
	teacherRepo := postgres.NewTeacherRepository(postgres.DB)
	return &TeacherService{
		TeacherRepository: teacherRepo,
	}
}

type Teacher interface {
	CreateTeacher(teacher *models.Teacher) error
	GetTeacherByID(id int) (*models.Teacher, error)
	ListTeachers(filter map[string]interface{}) ([]*models.Teacher, error)
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(id int) error
}

func (s *TeacherService) CreateTeacher(teacher *models.Teacher) error {
	return s.TeacherRepository.CreateTeacher(teacher)
}

func (s *TeacherService) GetTeacherByID(id int) (*models.Teacher, error) {
	return s.TeacherRepository.GetTeacherByID(id)
}

func (s *TeacherService) ListTeachers(filter map[string]interface{}) ([]*models.Teacher, error) {
	return s.TeacherRepository.ListTeachers(filter)
}

func (s *TeacherService) UpdateTeacher(teacher *models.Teacher) error {
	return s.TeacherRepository.UpdateTeacher(teacher)
}

func (s *TeacherService) DeleteTeacher(id int) error {
	return s.TeacherRepository.DeleteTeacher(id)
}
