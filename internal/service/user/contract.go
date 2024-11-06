package user

import (
	"context"
	"homeworktodolist/internal/entity"
)

type UserRepo interface {
	Create(ctx context.Context, user entity.User) error
}
