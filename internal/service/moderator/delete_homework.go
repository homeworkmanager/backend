package moderator

import (
	"context"

	"homeworktodolist/internal/entity"
)

func (s *Service) DeleteHomework(ctx context.Context, id entity.HomeworkID) error {
	err := s.homeworkService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
