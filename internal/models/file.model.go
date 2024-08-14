package models

import "time"

type File struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	FileName    string    `gorm:"size:255;not null"`
	URL         string    `gorm:"size:255;not null"`
	Grade       int64     `gorm:"not null"`
	Description string    `gorm:"size:255"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
