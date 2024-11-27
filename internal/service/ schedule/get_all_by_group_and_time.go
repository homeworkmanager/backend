package schedule

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

type GetSchedule struct {
	groupId  entity.GroupID
	fromTime time.Time
	toTime   time.Time
}

func (s *Service) GetAllByGroupAndTime(ctx context.Context, req GetSchedule) ([]entity.Class, error) {
	return []entity.Class{}, nil
}
