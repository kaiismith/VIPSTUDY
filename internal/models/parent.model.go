package models

import "time"

type Parent struct {
	ID              int64     `gorm:"primaryKey;autoIncrement"`
	ParentFirstName string    `gorm:"size:50;not null"`
	ParentLastName  string    `gorm:"size:50;not null"`
	PhoneNumber     string    `gorm:"size:20;not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
