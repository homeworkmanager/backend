package group

import (
	"context"
	"homeworktodolist/internal/entity"
)

type GroupRepo interface {
	Create(ctx context.Context, group entity.Group) (entity.GroupID, error)
	GetByName(ctx context.Context, name string) (entity.Group, error)
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}
