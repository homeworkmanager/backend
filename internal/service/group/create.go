package group

import (
	"context"
	"errors"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

type CreateGroup struct {
	Name     string
	Course   int8
	IcalLink string
}

func (s *Service) Create(ctx context.Context, req CreateGroup) (entity.GroupID, error) {
	group := req.toGroup()

	_, err := s.groupRepo.GetByName(ctx, group.Name)

	if err == nil {
		return 0, errs.GroupExists
	}
	if !(errors.Is(err, errs.GroupNotFound)) {
		return 0, err
	}

	id, err := s.groupRepo.Create(ctx, group)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (g *CreateGroup) toGroup() entity.Group {
	return entity.Group{
		Name:     g.Name,
		Course:   g.Course,
		IcalLink: g.IcalLink,
	}
}
