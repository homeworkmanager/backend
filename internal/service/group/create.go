package group

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

func (s *Service) Create(ctx context.Context, group entity.Group) (entity.GroupID, error) {

	_, err := s.groupRepo.GetByName(ctx, group.Name)

	if err == nil {
		return 0, errs.GroupExists
	}
	if !(errors.Is(err, errs.GroupNotFound)) {
		return 0, err
	}

	group.RegisterKey = uuid.NewString()[:5]

	id, err := s.groupRepo.Create(ctx, group)
	if err != nil {
		return 0, err
	}

	return id, nil

}
