package group

import (
	"context"

	"homeworktodolist/internal/entity"
)

func (s *Service) GetAllGroups(ctx context.Context) ([]entity.Group, error) {
	groups, err := s.groupRepo.GetAllGroups(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
