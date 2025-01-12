package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) Update(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error {
	err := s.homeworkService.Update(ctx, homeworkId, homeworkText)
	if err != nil {
		return err
	}
	return nil
}
