package admin

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"

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

func MapAdminRoutes(g fiber.Router, h *Handler, mw *middleware.MwManager) {
	g.Post("/addGroup", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin}), h.AddGroup())
	g.Patch("/updateClasses", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin}), h.UpdateClasses())
	g.Patch("/refreshAllData", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin}), h.RefreshAllData())
	g.Patch("/role/:userID", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin}), h.UpdateRole())
	g.Get("/users", mw.Auth(), mw.AllowedRoles([]entity.Role{entity.RoleGlobalAdmin}), h.GetAllUsers())
}
