package services

import "github.com/PedroWiller/all-day/internal/domain/repositories"

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}
