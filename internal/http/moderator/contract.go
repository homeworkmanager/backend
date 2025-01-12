package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
	moderatorService "homeworktodolist/internal/service/moderator"
)

type ModeratorService interface {
	AddHomework(ctx context.Context, req moderatorService.AddHomework) (entity.HomeworkID, error)
	DeleteHomework(ctx context.Context, id entity.HomeworkID) error
	Update(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error
}
