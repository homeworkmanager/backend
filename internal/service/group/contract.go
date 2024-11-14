package group

import (
	"context"
	"homeworktodolist/internal/entity"
)

type GroupRepo interface {
	Create(ctx context.Context, group entity.Group) error
	GetByName(ctx context.Context, name string) (entity.Group, error)
}
