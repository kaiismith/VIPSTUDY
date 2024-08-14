package models

import "time"

type Log struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	Action       string    `gorm:"size:255;not null"`
	IsAttendance bool      `gorm:"not null"`
	UserID       int64     `gorm:"foreignKey:UserID;references:ID;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
