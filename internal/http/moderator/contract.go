package moderator

import (
	"context"
	moderatorService "homeworktodolist/internal/service/moderator"
)

type ModeratorService interface {
	AddGroup(ctx context.Context, req moderatorService.AddHomework) error
}
