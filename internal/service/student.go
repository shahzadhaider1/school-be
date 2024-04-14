package service

import (
	"school-be/internal/repository/postgres"
	"school-be/pkg/models"
)

type StudentService struct {
	StudentRepository postgres.StudentRepository
}

func NewStudentService() *StudentService {
	studentRepo := postgres.NewStudentRepository(postgres.DB)
	return &StudentService{
		StudentRepository: studentRepo,
	}
}

type Student interface {
	CreateStudent(Student *models.Student) error
	GetStudentByID(id int) (*models.Student, error)
	ListStudents(filter map[string]interface{}) ([]*models.Student, error)
	UpdateStudent(Store *models.Student) error
	DeleteStudent(id int) error
}

func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.StudentRepository.CreateStudent(student)
}

func (s *StudentService) GetStudentByID(id int) (*models.Student, error) {
	return s.StudentRepository.GetStudentByID(id)
}

func (s *StudentService) ListStudents(filter map[string]interface{}) ([]*models.Student, error) {
	return s.StudentRepository.ListStudents(filter)
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	return s.StudentRepository.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.StudentRepository.DeleteStudent(id)
}
