package subject

import (
	"context"
	"homeworktodolist/internal/entity"
)

type SubjectService interface {
	GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error)
}
