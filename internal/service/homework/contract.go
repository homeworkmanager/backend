package homework

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

type HomeworkRepo interface {
	Clear(ctx context.Context) error
	Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error)
	Delete(ctx context.Context, id entity.HomeworkID) error
	Update(ctx context.Context, id entity.HomeworkID, homeworkText string) error
	GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error)
}
