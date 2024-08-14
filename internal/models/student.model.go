package models

import "time"

type Student struct {
	ID               int64     `gorm:"primaryKey;autoIncrement"`
	StudentFirstName string    `gorm:"size:50;not null"`
	StudentLastName  string    `gorm:"size:50;not null"`
	DateOfBirth      time.Time `gorm:"type:date"`
	School           string    `gorm:"size:100;not null"`
	Grade            int64     `gorm:"not null"`
	EntryMark        float64   `gorm:"not null"`
	PhoneNumber      string    `gorm:"size:20"`
	IsRemote         bool      `gorm:"default:false;not null"`
	Status           string    `gorm:"size:50;not null"`
	ParentID         int64     `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DesiredSchool    string    `gorm:"size:100;not null"`
	DesiredMark      float64   `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
}
