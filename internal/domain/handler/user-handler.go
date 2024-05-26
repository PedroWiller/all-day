package handler

import "github.com/PedroWiller/all-day/internal/domain/services"

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
