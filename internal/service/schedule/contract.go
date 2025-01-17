package schedule

import (
	"context"
	"time"

	"homeworktodolist/internal/entity"
)

type HomeworkService interface {
	GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error)
}

type ClassService interface {
	GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Class, error)
}
