package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ini/pkg/api/models"
	"ini/pkg/api/repositories"
)

var UserService userServiceI = new(userService)

type userServiceI interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
	GetUserByUsernamAndPassword(username, password string) (models.User, error)
}

type userService struct {
}

func (u userService) GetUsers() ([]models.User, error) {
	return repositories.UserRepo.GetUsers()
}

func (u userService) GetUser(id string) (models.User, error) {
	return repositories.UserRepo.GetUser(id)
}

func (u userService) CreateUser(user models.User) (models.User, error) {
	user.Id = primitive.NewObjectID()
	return repositories.UserRepo.CreateUser(user)
}

func (u userService) UpdateUser(user models.User) (models.User, error) {
	return repositories.UserRepo.UpdateUser(user)
}

func (u userService) DeleteUser(id string) error {
	return repositories.UserRepo.DeleteUser(id)
}

func (u userService) GetUserByUsernamAndPassword(username, password string) (models.User, error) {
	return repositories.UserRepo.GetUserByUsernameAndPassword(username, password)
}
