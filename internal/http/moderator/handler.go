package moderator

import "github.com/gofiber/fiber/v2"

type Handler struct {
	moderatorService ModeratorService
}

func NewModeratorHandler(moderatorService ModeratorService) *Handler {
	return &Handler{
		moderatorService: moderatorService,
	}
}

// TODO: проверка ролей
func MapModeratorRoutes(g fiber.Router, h *Handler) {
	g.Post("/addHomework/class", h.AddHomeworkToClass())
	g.Post("/addHomework/date", h.AddHomeworkToDate())
}
