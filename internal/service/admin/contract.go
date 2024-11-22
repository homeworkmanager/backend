package admin

import (
	"context"
	"homeworktodolist/internal/entity"
)

type GroupService interface {
	Create(ctx context.Context, group entity.Group) (entity.GroupID, error)
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}

type ClassService interface {
	UpdGroupClasses(ctx context.Context, g entity.Group) error
	ClearAllClasses(ctx context.Context) error
}

type SubjectService interface {
	UpdGroupSubjects(ctx context.Context, group entity.Group) error
	ClearAllSubjects(ctx context.Context) error
}

type HomeworkService interface {
	ClearAllHomeworks(ctx context.Context) error
}
