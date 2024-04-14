package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName   string `gorm:"not null" json:"first_name" validate:"required"`
	LastName    string `gorm:"not null" json:"last_name" validate:"required"`
	Address     string `gorm:"not null" json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	ClassID     uint   // many-to-one realtionship with Class
	// the default foreign key’s name is the owner’s type name plus the name of
	// its primary key field, so in this case: ClassID here is the foreign key
	Teachers []*Teacher `gorm:"many2many:student_teachers;"`
}
