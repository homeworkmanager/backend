package homework

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

type HomeworkRepo interface {
	Clear(ctx context.Context) error
	Create(ctx context.Context, homework entity.Homework) error
	GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error)
}
