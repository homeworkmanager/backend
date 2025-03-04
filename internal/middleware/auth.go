package middleware

import (
	"github.com/gofiber/fiber/v2"

	"homeworktodolist/internal/entity"
)

func (mw *MwManager) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies(entity.SessionKey)

		if sessionKey == "" {
			return fiber.ErrUnauthorized
		}

		userCreds, err := mw.userRedisRepo.GetCreds(c.Context(), sessionKey)
		if err != nil {
			return err
		}
		userCreds.Role, err = mw.lookupRole(c.Context(), userCreds.UserID)
		if err != nil {
			return err
		}
		c.Locals(entity.Claims, userCreds)

		return c.Next()

	}
}
