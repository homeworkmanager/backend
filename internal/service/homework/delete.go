package homework

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) Delete(ctx context.Context, id entity.HomeworkID) error {
	err := s.homeworkRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
