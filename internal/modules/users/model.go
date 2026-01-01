package users

import "time"

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"size:20;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
