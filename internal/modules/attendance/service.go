package attendance

import (
	"time"

	serr "github.com/gorunriki/akademiflow/internal/shared/errors"
)

type Service interface {
	CreateAttendance(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

const (
	AttendancePresent = "present"
)

// create attendance
func (s *service) CreateAttendance(userID uint) error {
	if userID == 0 {
		return serr.ErrInvalidInput
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	exists, err := s.repo.ExistsByUserAndDate(userID, today)
	if err != nil {
		return err
	}
	if exists {
		return serr.ErrConflict
	}

	attendance := Attendance{
		UserID: userID,
		Date:   today,
		Status: AttendancePresent,
	}

	return s.repo.Create(&attendance)
}
