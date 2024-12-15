package schedule

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

type HomeworkService interface {
	GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error)
}

type ClassService interface {
	GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Class, error)
}
