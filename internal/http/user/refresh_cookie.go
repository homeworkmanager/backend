package user

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	"time"
)

func (h *Handler) Refresh() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies(entity.SessionKey)

		if sessionKey == "" {
			return fiber.ErrForbidden
		}

		err := h.userService.RefreshCookie(c.Context(), sessionKey)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		expiresTime := time.Now().Add(h.config.AuthTTL)

		c.Cookie(&fiber.Cookie{
			Name:    entity.SessionKey,
			Value:   sessionKey,
			Path:    "/",
			Domain:  h.config.Domain,
			Expires: expiresTime,
		})
		c.Cookie(&fiber.Cookie{
			Name:    "session_expires",
			Value:   expiresTime.String(),
			Path:    "/",
			Domain:  h.config.Domain,
			Expires: expiresTime,
		})

		return c.JSON(fiber.Map{
			"data": "success",
		})

	}
}
