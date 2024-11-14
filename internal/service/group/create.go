package group

import (
	"context"
	"errors"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

type CreateGroup struct {
	Name   string
	Course int8
}

func (s *Service) Create(ctx context.Context, req CreateGroup) error {
	group := req.toGroup()

	_, err := s.groupRepo.GetByName(ctx, group.Name)

	if err == nil {
		return errs.GroupExists
	}
	if !(errors.Is(err, errs.GroupNotFound)) {
		return err
	}

	if err := s.groupRepo.Create(ctx, group); err != nil {
		return err
	}

	return nil

}

func (g *CreateGroup) toGroup() entity.Group {
	return entity.Group{
		Name:   g.Name,
		Course: g.Course,
	}
}
