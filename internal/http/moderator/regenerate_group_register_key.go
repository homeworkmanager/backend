package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
)

func (h *Handler) RegenerateGroupRegisterKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		registerKey, err := h.moderatorService.RegenerateGroupRegisterKey(c.Context(), creds.GroupID)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"register_key": registerKey,
			"data":         "Key successfully regenerated",
		})

	}
}
