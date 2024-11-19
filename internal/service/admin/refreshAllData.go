package admin

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/utils/classParse"
)

// TODO: дописать метод сделать транзакцию и удаление предыдущих
func (s *Service) RefreshAllData(ctx context.Context) error {

	groups, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}
	for _, group := range groups {
		class, err := classParse.IcalParse(group.IcalLink)
		if err != nil {
			return err
		}
		subjects := classParse.ParseSubject(class)
		for _, subject := range subjects {
			err = s.subjectService.Create(ctx, entity.Subject{
				GroupId:     group.GroupID,
				SubjectName: subject,
			})
			if err != nil {
				return err
			}
		}

	}
	return nil
}
