package database

import (
	"log"

	"github.com/gorunriki/akademiflow/internal/modules/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&users.User{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")
}
