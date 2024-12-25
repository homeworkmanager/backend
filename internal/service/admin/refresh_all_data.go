package admin

import (
	"context"
	"homeworktodolist/internal/utils/classParse"
)

// TODO: проверить что функция err = s.homeworkService.ClearAllHomeworks(ctx) правильно отрабатывает
func (s *Service) RefreshAllData(ctx context.Context) error {

	groups, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	err = s.manager.Do(ctx, func(ctx context.Context) error {
		err = s.homeworkService.ClearAllHomeworks(ctx)
		if err != nil {
			return err
		}
		err = s.classService.ClearAllClasses(ctx)
		if err != nil {
			return err
		}

		err = s.subjectService.ClearAllSubjects(ctx)
		if err != nil {
			return err
		}

		for _, group := range groups {

			classes, subjects, err := classParse.IcalParse(group.IcalLink)
			if err != nil {
				return err
			}

			err = s.subjectService.UpdGroupSubjects(ctx, group, subjects)
			if err != nil {
				return err
			}

			err = s.classService.UpdGroupClasses(ctx, group, classes)
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
