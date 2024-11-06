package user

import "homeworktodolist/internal/config"

type Handler struct {
	config *config.Config
	userService UserService
}

func NewUserHandler(config *config.Config, service UserService) *Handler {
	return &Handler{
		config:      config,
		userService: service,
	}
}