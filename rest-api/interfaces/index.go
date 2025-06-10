package interfaces

import "learning-golang/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById() (models.User, error)
	CreateUser(models.User) (int64, error)
	Login(models.User) (string, error)
}
