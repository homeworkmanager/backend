package class

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/utils/classParse"
)

type UpdateGroup struct {
	GroupId  entity.GroupID
	IcalLink string
	Class    []entity.UpdateClass
}

func (s *Service) UpdGroupClasses(ctx context.Context, g entity.Group) error {
	var err error

	group := toUpdateGroup(g)

	group.Class, err = classParse.IcalParse(group.IcalLink)

	if err != nil {
		return err
	}

	for i := range group.Class {
		subject, err := s.subjectService.GetBySubjectNameAndGroup(ctx, string([]rune(group.Class[i].Summary)[3:]), group.GroupId)
		if err != nil {
			return err
		}
		group.Class[i].SubjectID = subject.SubjectId
	}

	classes := generateGroupSchedule(entity.SemesterNumWeeks, group)
	err = s.classRepo.Create(ctx, classes)
	if err != nil {
		return err
	}

	return nil

}

func toUpdateGroup(group entity.Group) UpdateGroup {
	return UpdateGroup{
		GroupId:  group.GroupID,
		IcalLink: group.IcalLink,
		Class:    []entity.UpdateClass{},
	}
}
