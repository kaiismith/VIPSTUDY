package models

import "time"

type Exam struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Mark      float64   `gorm:"not null"`
	StudentID int64     `gorm:"foreignKey:StudentID;references:ID;not null"`
	ClassID   int64     `gorm:"foreignKey:ClassID;references:ID;not null"`
	FileID    int64     `gorm:"foreignKey:FileID;references:ID;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
