package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name       string    `gorm:"unique;not null" json:"name" validate:"required"`
	RoomNumber string    `gorm:"unique;not null" json:"room_number" validate:"required"`
	Students   []Student // one-to-many relationship with Student
	// the default foreign key’s name is the owner’s type name plus the name of
	// its primary key field, so in this case: ClassID
	// To use another field as foreign key, you can customize it with a foreignKey tag,
	// e.g: `gorm:"foreignKey:ClassIDNew"`
	// and instead of "ClassID", this "ClassIDNew" should be present in Student Model.
	// Students   []Student `gorm:"foreignKey:ClassIDNew"`
	Teachers []*Teacher `gorm:"many2many:class_teachers;"`
}
