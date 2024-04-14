package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FirstName   string     `gorm:"not null" json:"first_name" validate:"required"`
	LastName    string     `gorm:"not null" json:"last_name" validate:"required"`
	Address     string     `gorm:"not null" json:"address" validate:"required"`
	PhoneNumber string     `json:"phone_number" validate:"required"`
	Students    []*Student `gorm:"many2many:student_teachers;"`
	Classes     []*Class   `gorm:"many2many:class_teachers;"`
}
