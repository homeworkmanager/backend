package group

import (
	"context"

	"homeworktodolist/internal/entity"
)

type GroupService interface {
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}
