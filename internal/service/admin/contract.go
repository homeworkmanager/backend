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
	UpdGroupClasses(ctx context.Context, g entity.Group) error
	ClearAllClasses(ctx context.Context) error
}

type SubjectService interface {
	UpdGroupSubjects(ctx context.Context, group entity.Group) error
}
