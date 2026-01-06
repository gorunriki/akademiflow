package auth

import (
	"github.com/gorunriki/akademiflow/internal/modules/users"
	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*users.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (*users.User, error) {
	var user users.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
