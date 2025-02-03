package middleware

import (
	"github.com/gofiber/fiber/v2"

	"homeworktodolist/internal/entity"
)

func (mw *MwManager) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies(entity.SessionKey)

		if sessionKey == "" {
			return fiber.ErrForbidden
		}

		userCreds, err := mw.UserRedisRepo.GetCreds(c.Context(), sessionKey)
		if err != nil {
			return err
		}
		c.Locals(entity.Claims, userCreds)

		return c.Next()

	}
}
