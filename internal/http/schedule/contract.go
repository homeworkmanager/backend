package schedule

import (
	"context"

	"homeworktodolist/internal/entity"
	scheduleService "homeworktodolist/internal/service/schedule"
)

type ScheduleService interface {
	GetAllByGroupAndTime(ctx context.Context, req scheduleService.GetSchedule) ([]entity.ScheduleDay, error)
	GetHomeworks(ctx context.Context, req scheduleService.GetHomework) ([]entity.HomeworkDay, error)
}
