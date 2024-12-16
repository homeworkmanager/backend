package schedule

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/middleware"
)

type Handler struct {
	scheduleService ScheduleService
}

func NewScheduleHandler(scheduleService ScheduleService) *Handler {
	return &Handler{
		scheduleService: scheduleService,
	}
}

func MapScheduleRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Get("/get", mw.Auth(), h.GetSchedule())
}
