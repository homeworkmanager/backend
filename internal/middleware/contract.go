package middleware

import (
	"context"
	"homeworktodolist/internal/entity"
)

type UserRedisRepo interface {
	GetCreds(ctx context.Context, sessionKey string) (entity.UserCreds, error)
}
