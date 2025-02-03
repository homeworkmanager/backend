package entity

import "time"

type Homework struct {
	HomeworkID     HomeworkID
	SemClassNumber *int64
	GroupID        GroupID
	SubjectID      SubjectID
	HomeworkText   string
	Category       *ClassCategory
	IsCompleted    bool
	DueDate        time.Time
	CreatedAt      time.Time
}
