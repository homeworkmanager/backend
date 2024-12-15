package homework

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

func (s *Service) GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error) {
	homeworks, err := s.homeworkRepo.GetByGroupAndTime(ctx, groupID, fromTime, toTime)
	if err != nil {
		return nil, err
	}
	return homeworks, nil
}
