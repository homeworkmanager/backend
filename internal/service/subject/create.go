package subject

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) Create(ctx context.Context, subject entity.Subject) error {

	err := s.subjectRepo.Create(ctx, subject)

	if err != nil {
		return err
	}
	return nil
}
