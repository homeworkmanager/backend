package user

import (
	"context"
	userService "homeworktodolist/internal/service/user"
)

type UserService interface {
	Create(ctx context.Context, req userService.CreateUser) error
}