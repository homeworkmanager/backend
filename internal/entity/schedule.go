package entity

import "time"

type Schedule struct {
	ClassID        ClassID
	GroupID        GroupID
	SubjectID      SubjectID
	StartTime      time.Time
	EndTime        time.Time
	Summary        string
	SemClassNumber int64
	Location       string
	HomeworkID     HomeworkID
	HomeworkText   string
	DueDate        time.Time
}
