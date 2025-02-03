package moderator

import (
	"github.com/gofiber/fiber/v2"

	"homeworktodolist/internal/middleware"
)

type Handler struct {
	moderatorService ModeratorService
}

func NewModeratorHandler(moderatorService ModeratorService) *Handler {
	return &Handler{
		moderatorService: moderatorService,
	}
}

// TODO: проверка ролей
func MapModeratorRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/addHomework/class", mw.Auth(), h.AddHomeworkToClass())
	g.Post("/addHomework/date", mw.Auth(), h.AddHomeworkToDate())
	g.Delete("/delete", mw.Auth(), h.DeleteHomework())
	g.Patch("/update", mw.Auth(), h.UpdateHomework())
	g.Post("/note/add/:subjectID", mw.Auth(), h.AddNote())
	g.Delete("/note/delete/:noteID", mw.Auth(), h.DeleteNote())
	g.Patch("/note/update/:noteID", mw.Auth(), h.UpdateNote())
}
