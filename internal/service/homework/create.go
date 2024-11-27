package homework

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) Create(ctx context.Context, homework entity.Homework) error {

	err := s.homeworkRepo.Create(ctx, homework)
	if err != nil {
		return err
	}
	return nil
}
