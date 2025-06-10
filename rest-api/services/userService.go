package services

import (
	"learning-golang/interfaces"
	"learning-golang/models"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(models.User) (int64, error)
	Login(models.User) (string, error)
}

type userService struct {
	repo interfaces.UserRepository
}

// pass the repository dependency
func NewUserService(r interfaces.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) CreateUser(u models.User) (int64, error) {
	return s.repo.CreateUser(u)
}

func (s *userService) Login(u models.User) (string, error) {
	return s.repo.Login(u)
}
