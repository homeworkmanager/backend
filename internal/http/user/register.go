package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	userService "homeworktodolist/internal/service/user"
)

type RegisterReq struct {
	Name     string  `json:"name"`
	Surname  *string `json:"surname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	GroupId  int64   `json:"groupId"`
}

func (h *Handler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		mName := "UserHandler.Register"
		var req RegisterReq

		if err := c.BodyParser(&req); err != nil {
			err := errors.Wrap(err, fiber.ErrBadRequest.Error())
			return errors.Wrap(err, mName)
		}

		err := h.userService.Create(c.Context(), req.toCreate())
		if err != nil {
			return errors.Wrap(err, mName)
		}

		return c.JSON(fiber.Map{
			"data": "Successfully registered",
		})
	}
}

func (r *RegisterReq) toCreate() userService.CreateUser {
	return userService.CreateUser{
		Name:     r.Name,
		Surname:  r.Surname,
		Email:    r.Email,
		Password: r.Password,
		GroupId:  r.GroupId,
	}
}
