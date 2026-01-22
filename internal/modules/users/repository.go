package users

import "gorm.io/gorm"

type Repository interface {
	FindByID(id uint) (*User, error)
	Create(user *User) error
	ExistsByEmail(email string) (bool, error)
	ListUsers(limit int, offset int, keyword string, includeDeleted bool) ([]User, int64, error)
	Delete(id uint) error
	UpdateRole(id uint, role string) error
	Restore(id uint) error
	FindByIDIncludingDeleted(id uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// user details
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

// check user availability
func (r *repository) FindByIDIncludingDeleted(id uint) (*User, error) {
	var user User
	if err := r.db.Unscoped().First(&user, id).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

// list users
func (r *repository) ListUsers(limit int, offset int, keyword string, includeDeleted bool) ([]User, int64, error) {
	var users []User
	var total int64

	query := r.db.Model(&User{})

	if keyword != "" {
		like := "%" + keyword + "%"
		query = query.Where(
			"name ILIKE ? OR email ILIKE ?", like, like,
		)
	}

	if includeDeleted {
		query = query.Unscoped()
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Delete user
func (r *repository) Delete(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// update user role
func (r *repository) UpdateRole(id uint, role string) error {
	result := r.db.Model(&User{}).Where("id = ?", id).Update("role", role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Restore deleted user
func (r *repository) Restore(id uint) error {
	result := r.db.Unscoped().Model(&User{}).Where("id = ? AND deleted_at IS NOT NULL", id).Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
