package admin

import (
	"context"
	"homeworktodolist/internal/entity"
	groupService "homeworktodolist/internal/service/group"
)

type AddGroup struct {
	Name     string
	Course   int8
	IcalLink string
}

func (s *Service) AddGroup(ctx context.Context, req AddGroup) error {
	group := req.toUser()

	err := s.groupService.Create(ctx, groupService.CreateGroup{
		Name:     group.Name,
		Course:   group.Course,
		IcalLink: group.IcalLink,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *AddGroup) toUser() entity.Group {
	return entity.Group{
		Name:     r.Name,
		Course:   r.Course,
		IcalLink: r.IcalLink,
	}
}
