package homework

import (
	"context"
	"homeworktodolist/internal/entity"
)

type HomeworkRepo interface {
	Clear(ctx context.Context) error
	Create(ctx context.Context, homework entity.Homework) error
}
