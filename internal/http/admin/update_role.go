package admin

import (
	"github.com/gofiber/fiber/v2"
	adminService "homeworktodolist/internal/service/admin"
)

func (h *Handler) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_ = c.Params("id")

		err := h.adminService.UpdateRole(c.Context(), adminService.UpdateUserRole{
			UserID: 0,
			Role:   0,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "User role updated successfully",
		})
	}
}
