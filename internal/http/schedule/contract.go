package schedule

import (
	"context"

	"homeworktodolist/internal/entity"
	scheduleService "homeworktodolist/internal/service/schedule"
)

type ScheduleService interface {
	GetAllByGroupAndTime(ctx context.Context, req scheduleService.GetSchedule) ([]entity.Day, error)
}
