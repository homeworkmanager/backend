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
	Class    []UpdateClass
}
type UpdateClass struct {
	SubjectID   entity.SubjectID
	Summary     string
	Start       time.Time
	End         time.Time
	Description string
	Location    string
}

func (s *Service) Update(ctx context.Context) error {
	g, err := s.groupService.GetAllGroups(ctx)
	if err != nil {
		return err
	}

	groups := toUpdateGroups(g)

	for _, group := range groups {

		group.Class, err = classParse.IcalParse(group.IcalLink)
		if err != nil {
			return err
		}

		for _, class := range group.Class {
			subject, err := s.subjectService.GetBySubjectNameAndGroup(ctx, class.Summary[3:], group.GroupId)
			if err != nil {
				return err
			}
			class.SubjectID = subject.SubjectId
		}

	}

	return err
}

func toUpdateGroups(groups []entity.Group) []UpdateGroup {
	upd := make([]UpdateGroup, len(groups))
	for i, group := range groups {
		upd[i] = UpdateGroup{
			GroupId:  group.GroupID,
			IcalLink: group.IcalLink,
			Class:    []UpdateClass{},
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
