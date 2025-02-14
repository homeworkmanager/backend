package group

import (
	"context"
)

func (s *Service) RegenerateRegisterKeys(ctx context.Context) error {
	groups, err := s.groupRepo.GetAllGroups(ctx)
	if err != nil {
		return err
	}
	for i, _ := range groups {
		err = s.RegenerateGroupRegisterKey(ctx, groups[i].GroupID)
		if err != nil {
			return err
		}
	}
	return nil
}
