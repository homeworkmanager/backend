package user

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
)

type UserResp struct {
	UserID  entity.UserID  `json:"user_id"`
	Name    string         `json:"name"`
	Surname *string        `json:"surname"`
	Email   string         `json:"email"`
	Role    entity.Role    `json:"role"`
	GroupID entity.GroupID `json:"group_id"`
}

func (h Handler) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)

		if !ok {
			return fiber.ErrUnauthorized
		}

		user, err := h.userService.GetUser(c.Context(), creds.UserID)
		if err != nil {
			return err
		}

		return c.JSON(toUserResp(user))
	}
}

func toUserResp(user entity.User) UserResp {
	return UserResp{
		UserID:  user.UserID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Role:    user.Role,
		GroupID: user.GroupID,
	}
}
