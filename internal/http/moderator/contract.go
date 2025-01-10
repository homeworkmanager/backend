package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
	moderatorService "homeworktodolist/internal/service/moderator"
)

type ModeratorService interface {
	AddHomework(ctx context.Context, req moderatorService.AddHomework) (entity.HomeworkID, error)
}
