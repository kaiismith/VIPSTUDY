package models

import "time"

type Class struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	ClassName   string    `gorm:"size:50;not null"`
	Status      string    `gorm:"size:50;not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
