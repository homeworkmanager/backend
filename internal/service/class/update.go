package class

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/utils/classParse"
	"time"
)

type UpdateGroup struct {
	GroupId  entity.GroupID
	IcalLink string
	Class    []entity.UpdateClass
}

// TODO: Подумать на тем может стоит объеденить все в один sql запрос, а не по группам
func (s *Service) Update(ctx context.Context) error {
	g, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	groups := toUpdateGroups(g)

	err = s.manager.Do(ctx, func(ctx context.Context) error {
		err = s.classRepo.Clear(ctx)
		if err != nil {
			return err
		}
		for _, group := range groups {

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

			classes := generateFullGroupSchedule(16, group)
			err := s.classRepo.Create(ctx, classes)
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

func toUpdateGroups(groups []entity.Group) []UpdateGroup {
	upd := make([]UpdateGroup, len(groups))
	for i, group := range groups {
		upd[i] = UpdateGroup{
			GroupId:  group.GroupID,
			IcalLink: group.IcalLink,
			Class:    []entity.UpdateClass{},
		}
	}
	return upd
}

func generateFullGroupSchedule(weeks int, updateGroup UpdateGroup) []entity.Class {
	twoWeekSchedule := updateGroup.Class
	var fullSchedule []entity.Class

	subjectCounters := make(map[entity.SubjectID]int64)

	for i := 0; i < weeks/2; i++ {
		for _, class := range twoWeekSchedule {

			shift := time.Duration(i*14*24) * time.Hour

			subjectCounters[class.SubjectID]++

			fullSchedule = append(fullSchedule, entity.Class{
				GroupID:        updateGroup.GroupId,
				SubjectID:      class.SubjectID,
				StartTime:      class.Start.Add(shift),
				EndTime:        class.End.Add(shift),
				Summary:        class.Summary,
				SemClassNumber: subjectCounters[class.SubjectID],
				Location:       class.Location,
			})
		}
	}
	return fullSchedule
}
