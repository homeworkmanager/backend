package subject

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/utils/classParse"
)

func (s *Service) UpdGroupSubjects(ctx context.Context, group entity.Group) error {

	class, err := classParse.IcalParse(group.IcalLink)
	if err != nil {
		return err
	}
	subjects := classParse.ParseSubject(class)
	for _, subject := range subjects {
		err = s.subjectRepo.Create(ctx, entity.Subject{
			GroupId:     group.GroupID,
			SubjectName: subject,
		})
		if err != nil {
			return err
		}
	}
	return nil

}
