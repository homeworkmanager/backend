package class

import (
	"context"
	"homeworktodolist/internal/entity"
	groupService "homeworktodolist/internal/service/group"
)

type GroupService interface {
	Create(ctx context.Context, req groupService.CreateGroup) error
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}

type ClassRepo interface {
	Clear(ctx context.Context) error
}

type SubjectService interface {
	GetBySubjectNameAndGroup(ctx context.Context, subjectName string, groupId entity.GroupID) (entity.Subject, error)
}
