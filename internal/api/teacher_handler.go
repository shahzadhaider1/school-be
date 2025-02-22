package api

import (
	"school-be/internal/service"
)

type TeacherHandler struct {
	teacherService *service.TeacherService
}

func NewTeacherHandler() *TeacherHandler {
	return &TeacherHandler{}
}
