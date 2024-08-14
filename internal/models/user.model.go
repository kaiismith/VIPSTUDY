package models

import "time"

type User struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	Username      string    `gorm:"size:50;not null;unique"`
	Email         string    `gorm:"size:100;not null;unique"`
	Password      string    `gorm:"size:50;not null"`
	UserFirstName string    `gorm:"size:50;not null"`
	UserLastName  string    `gorm:"size:50;not null"`
	DateOfBirth   time.Time `gorm:"type:date"`
	PhoneNumber   string    `gorm:"size:20;not null"`
	University    string    `gorm:"size:100;not null"`
	QRCode        []byte    `gorm:"size:100"`
	IsAdmin       bool      `gorm:"default:false;not null"`
	IsMaintainer  bool      `gorm:"default:false;not null"`
	AccessToken   string    `gorm:"size:255;not null"`
	RefreshToken  string    `gorm:"size:255;not null"`
	Description   string    `gorm:"size:255"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
