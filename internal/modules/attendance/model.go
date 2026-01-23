package attendance

import (
	"time"

	"github.com/gorunriki/akademiflow/internal/modules/users"
)

type Attendance struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null;index"`
	Date      time.Time `gorm:"type:date;not null"`
	Status    string    `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User users.User `gorm:"foreignKey:UserID"`
}
