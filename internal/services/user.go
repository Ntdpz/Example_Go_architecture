package services

import (
	"Example_Go_architecture/internal/repositories"
	"Example_Go_architecture/models"

)

type UserService interface {
	GetUsers() ([]models.Users, error)
	GetUsersByParams(username, token string) ([]models.Users, error)
}
type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}
func (s *userService) GetUsers() ([]models.Users, error) {
	return s.repo.GetAllUsers()
}
func (s *userService) GetUsersByParams(username, token string) ([]models.Users, error) {
	return s.repo.FindUserByParams(username, token)
}
