package entity

import "time"

type Class struct {
	ClassID        ClassID
	GroupID        GroupID
	SubjectID      SubjectID
	StartTime      time.Time
	EndTime        time.Time
	Summary        string
	SemClassNumber int64
	Location       string
}

// TODO: спросить что делать, потому что если эта структура будет в сервисном слое будут циклические зависисмости
type UpdateClass struct {
	SubjectID   SubjectID
	Summary     string
	Start       time.Time
	End         time.Time
	Description string
	Location    string
}
