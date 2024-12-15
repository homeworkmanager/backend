package class

import (
	"homeworktodolist/internal/entity"
	"time"
)

type class struct {
	ClassID        entity.ClassID   `db:"class_id"`
	GroupID        entity.GroupID   `db:"group_id"`
	SubjectID      entity.SubjectID `db:"subject_id"`
	StartTime      time.Time        `db:"start_time"`
	EndTime        time.Time        `db:"end_time"`
	Summary        string           `db:"summary"`
	SemClassNumber int64            `db:"class_sem_number"`
	Location       string           `db:"location"`
}

func (c class) toClass() entity.Class {
	return entity.Class{
		ClassID:        c.ClassID,
		GroupID:        c.GroupID,
		SubjectID:      c.SubjectID,
		StartTime:      c.StartTime,
		EndTime:        c.EndTime,
		Summary:        c.Summary,
		SemClassNumber: c.SemClassNumber,
		Location:       c.Location,
	}
}
