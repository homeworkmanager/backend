package class

import (
	"homeworktodolist/internal/entity"
	"time"
)

func generateGroupSchedule(weeks int, updateGroup UpdateGroup) []entity.Class {
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
				Description:    class.Description,
				SemClassNumber: subjectCounters[class.SubjectID],
				Location:       class.Location,
			})
		}
	}
	return fullSchedule
}
