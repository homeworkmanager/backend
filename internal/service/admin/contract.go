package admin

import (
	"context"
	groupService "homeworktodolist/internal/service/group"
)

type GroupService interface {
	Create(ctx context.Context, req groupService.CreateGroup) error
}
