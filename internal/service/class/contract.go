package class

import (
	"context"
	"homeworktodolist/internal/entity"
)

type GroupService interface {
	GetAllGroups(ctx context.Context) ([]entity.Group, error)
}

type ClassRepo interface {
	Clear(ctx context.Context) error
	Create(ctx context.Context, classes []entity.Class) error
}

type SubjectService interface {
	GetBySubjectNameAndGroup(ctx context.Context, subjectName string, groupId entity.GroupID) (entity.Subject, error)
}
