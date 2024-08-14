package models

import "time"

type Attendance struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	StudentID   int64     `gorm:"foreignKey:StudentID;references:ID;not null"`
	FileID      int64     `gorm:"foreignKey:FileID;references:ID;not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
