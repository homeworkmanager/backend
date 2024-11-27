package schedule

import (
	"homeworktodolist/internal/entity"
	"time"
)

type schedule struct {
	ClassID        *entity.ClassID    `db:"class_id"`
	GroupID        entity.GroupID     `db:"group_id"`
	SubjectID      entity.SubjectID   `db:"subject_id"`
	StartTime      *time.Time         `db:"start_time"`
	EndTime        *time.Time         `db:"end_time"`
	Summary        *string            `db:"summary"`
	SemClassNumber *int64             `db:"class_sem_number"`
	Location       *string            `db:"location"`
	HomeworkID     *entity.HomeworkID `db:"homework_id"`
	HomeworkText   *string            `db:"homework_text"`
	DueDate        *time.Time         `db:"due_date"`
}

func (c schedule) toSchedule() entity.Schedule {
	return entity.Schedule{
		ClassID:        *c.ClassID,
		GroupID:        c.GroupID,
		SubjectID:      c.SubjectID,
		StartTime:      *c.StartTime,
		EndTime:        *c.EndTime,
		Summary:        *c.Summary,
		SemClassNumber: *c.SemClassNumber,
		Location:       *c.Location,
		HomeworkID:     *c.HomeworkID,
		HomeworkText:   *c.HomeworkText,
		DueDate:        *c.DueDate,
	}
}

func toSchedules(schedules []schedule) []entity.Schedule {
	res := make([]entity.Schedule, len(schedules))
	for i, g := range schedules {
		res[i] = g.toSchedule()
	}
	return res
}
