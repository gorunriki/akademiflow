package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetMe(userID uint) (*MeReponse, error)
	CreateUser(user *User) error
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
var ErrEmailAlreadyExists = errors.New("email already exists")

func (s *service) CreateUser(user *User) error {
	exists, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return ErrEmailAlreadyExists
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hasedPassword)
	user.Role = "user"

	return s.repo.Create(user)
}
