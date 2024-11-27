package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

type HomeworkService interface {
	Create(ctx context.Context, homework entity.Homework) error
}
