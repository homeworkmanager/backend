package moderator

import (
	"context"
	moderatorService "homeworktodolist/internal/service/moderator"
)

type ModeratorService interface {
	AddHomework(ctx context.Context, req moderatorService.AddHomework) error
}
