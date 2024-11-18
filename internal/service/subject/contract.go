package subject

import (
	"context"
	"homeworktodolist/internal/entity"
)

type SubjectRepo interface {
	GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error)
	GetByNameAndGroup(ctx context.Context, name string, groupID entity.GroupID) (entity.Subject, error)
}
