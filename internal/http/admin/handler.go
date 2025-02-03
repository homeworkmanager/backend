package admin

import (
	"github.com/gofiber/fiber/v2"

	"homeworktodolist/internal/middleware"
)

type Handler struct {
	adminService AdminService
}

func NewAdminHandler(service AdminService) *Handler {
	return &Handler{
		adminService: service,
	}
}

// TODO добавить валидацию user на то что он Admin
func MapAdminRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/addGroup", h.AddGroup())
	g.Patch("/updateClasses", h.UpdateClasses())
	g.Patch("/refreshAllData", h.RefreshAllData())
	g.Patch("/role/:userID", h.UpdateRole())
	g.Get("/users", h.GetAllUsers())
}
