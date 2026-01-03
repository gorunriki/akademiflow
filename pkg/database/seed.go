package database

import (
	"log"

	"github.com/gorunriki/akademiflow/internal/modules/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedAdmin(db)
}

func seedAdmin(db *gorm.DB) {
	const (
		adminEmail    = "admin@school.test"
		adminPassword = "admin123"
		adminRole     = "admin"
	)

	var count int64
	db.Model(&users.User{}).Where("email = ?", adminEmail).Count(&count)

	if count > 0 {
		log.Println("Admin user already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash admin password: ", err)
	}

	admin := users.User{
		Name:     "Admin",
		Email:    adminEmail,
		Password: string(hashedPassword),
		Role:     adminRole,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("failed to seed admin user: ", err)
	}

	log.Println("Admin user seeded successfully")
}
