package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

type HomeworkService interface {
	Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error)
	Delete(ctx context.Context, id entity.HomeworkID) error
	Update(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error
}
