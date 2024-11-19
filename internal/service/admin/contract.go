package admin

import (
	"context"
	"homeworktodolist/internal/entity"
	groupService "homeworktodolist/internal/service/group"
)

type GroupService interface {
	Create(ctx context.Context, req groupService.CreateGroup) error
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}

type ClassService interface {
	Update(ctx context.Context) error
}

type SubjectService interface {
	Create(ctx context.Context, subject entity.Subject) error
}
