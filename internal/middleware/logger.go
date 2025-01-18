package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (mw *MwManager) RequestLogger(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.Infow("Incoming request",
			"method", c.Method(),
			"url", c.OriginalURL(),
			"remote_ip", c.IP(),
		)

		return c.Next()
	}
}
