package users

import "gorm.io/gorm"

type Repository interface {
	FindByID(id uint) (*User, error)
	Create(user *User) error
	ExistsByEmail(email string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// create new user
func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
