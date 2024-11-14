package entity

import "time"

type Class struct {
	ClassID        ClassID
	GroupID        GroupID
	SubjectID      SubjectID
	Date           time.Time
	DayClassNumber int8
	SemClassNumber int64
	Location       string
}
