package admin

import (
	"context"
)

func (s *Service) UpdateClasses(ctx context.Context) error {

	g, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	err = s.manager.Do(ctx, func(ctx context.Context) error {
		err = s.classService.ClearAllClasses(ctx)
		if err != nil {
			return err
		}
		for _, group := range g {
			err = s.classService.UpdGroupClasses(ctx, group)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
