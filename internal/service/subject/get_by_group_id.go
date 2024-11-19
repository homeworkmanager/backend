package subject

import (
	"context"
	"homeworktodolist/internal/entity"
)

// TODO: Удалить если нахуй не нужна
func (s *Service) GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Subject, error) {
	subjects, err := s.subjectRepo.GetByGroupId(ctx, groupId)
	if err != nil {
		return subjects, err
	}
	return subjects, nil
}
