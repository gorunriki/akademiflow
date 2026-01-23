package database

import (
	"log"

	"github.com/gorunriki/akademiflow/internal/modules/attendance"
	"github.com/gorunriki/akademiflow/internal/modules/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&users.User{},
		&attendance.Attendance{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	db.Exec(`
			CREATE UNIQUE INDEX IF NOT EXISTS idx_attendances_user_date
			ON attendances (user_id, date)
		`)

	log.Println("Database migrated successfully")
}
