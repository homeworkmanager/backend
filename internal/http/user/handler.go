package user

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/config"
)

type Handler struct {
	config      *config.Config
	userService UserService
}

func NewUserHandler(config *config.Config, service UserService) *Handler {
	return &Handler{
		config:      config,
		userService: service,
	}
}

func MapUserRoutes(g fiber.Router, h *Handler) {
	g.Post("/register", h.Register())
	g.Post("/auth", h.Auth())
}
