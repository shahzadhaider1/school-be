package api

import "school-be/internal/service"

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler() *StudentHandler {
	return &StudentHandler{}
}
