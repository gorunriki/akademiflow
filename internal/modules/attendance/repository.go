package attendance

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(attendance *Attendance) error
	ExistsByUserAndDate(userID uint, date time.Time) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// check if user already create attendance today
func (r *repository) ExistsByUserAndDate(userID uint, date time.Time) (bool, error) {
	var count int64
	if err := r.db.Model(&Attendance{}).Where("user_id = ? AND date = ?", userID, date).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// present attendace
func (r *repository) Create(attendace *Attendance) error {
	if err := r.db.Create(attendace).Error; err != nil {
		return err
	}
	return nil
}
