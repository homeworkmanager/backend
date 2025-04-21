package entity

import "time"

type HomeworkFile struct {
	FileID     FileID
	HomeworkID HomeworkID
	FileName   string
	FileURL    string
	CreatedAt  time.Time
}
