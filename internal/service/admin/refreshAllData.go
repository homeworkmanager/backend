package admin

import (
	"context"
)

// TODO: дописать метод сделать транзакцию и удаление предыдущих
func (s *Service) RefreshAllData(ctx context.Context) error {

	groups, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	for _, group := range groups {
		err = s.subjectService.UpdGroupSubjects(ctx, group)
		if err != nil {
			return err
		}
	}

	return nil
}
