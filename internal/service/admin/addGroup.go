package admin

import (
	"context"
	"homeworktodolist/internal/entity"
	groupService "homeworktodolist/internal/service/group"
)

type AddGroup struct {
	Name   string
	Course int8
}

func (s *Service) AddGroup(ctx context.Context, req AddGroup) error {
	group := req.toUser()

	//TODO спросить у олега, нужно ли здесь переводить тип в CreateGroup из groupService, или можно сделать так чтобы функция сразу entity принимала
	err := s.groupService.Create(ctx, groupService.CreateGroup{
		Name:   group.Name,
		Course: group.Course,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *AddGroup) toUser() entity.Group {
	return entity.Group{
		Name:   r.Name,
		Course: r.Course,
	}
}
