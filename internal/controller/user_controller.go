package controller

import "github.com/PedroWiller/all-day/internal/domain/services"

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}
