package schedule

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewScheculeHandler() *Handler {
	return &Handler{}
}

func MapScheculeRoutes(g fiber.Router, h *Handler) {
}
