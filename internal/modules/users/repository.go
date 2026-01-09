package users

import "gorm.io/gorm"

type Repository interface {
	FindByID(id uint) (*User, error)
	Create(user *User) error
	ExistsByEmail(email string) (bool, error)
	ListUsers(limit, offset int) ([]User, int64, error)
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

// check email duplication
func (r *repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// list users
func (r *repository) ListUsers(limit, offset int) ([]User, int64, error) {
	var users []User
	var total int64

	if err := r.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Order("id ASC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
