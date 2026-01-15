package users

import (
	"errors"

	serr "github.com/gorunriki/akademiflow/internal/shared/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	GetMe(userID uint) (*MeReponse, error)
	CreateUser(user *User) error
	GetUsers(page int, limit int, keyword string) ([]UserResponse, int64, error)
	GetUser(id uint) (*UserResponse, error)
	DeleteUser(id uint) error
	UpdateUserRole(targetID uint, newRole string, currentUserID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetMe(userID uint) (*MeReponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &MeReponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

// create new user
// TODO 1.HASH PASSWORD, 2.SET ROLE DEFAULT, 3.PANGGIL repo.CreateUser(user)
func (s *service) CreateUser(user *User) error {
	exists, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return serr.ErrConflict
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	user.Role = "user"

	return s.repo.Create(user)
}

// get all user
func (s *service) GetUsers(page int, limit int, keyword string) ([]UserResponse, int64, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	users, total, err := s.repo.ListUsers(limit, offset, keyword)
	if err != nil {
		return nil, 0, err
	}

	res := make([]UserResponse, 0, len(users))
	for _, user := range users {
		res = append(res, UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		})
	}

	return res, total, nil
}

// user details
func (s *service) GetUser(id uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

// Delete user
func (s *service) DeleteUser(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serr.ErrNotFound
		}
		return err
	}
	return nil
}

// Update user role
func (s *service) UpdateUserRole(targetID uint, newRole string, currentUserID uint) error {
	if newRole != "admin" && newRole != "user" {
		return serr.ErrInvalidInput
	}
	if targetID == currentUserID && newRole == "user" {
		return serr.ErrForbidden
	}
	if err := s.repo.UpdateRole(targetID, newRole); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serr.ErrNotFound
		}
		return err
	}
	return nil
}
