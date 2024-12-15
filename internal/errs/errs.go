package errs

import "errors"

var (
	UserNotFound      = errors.New("user doesn't exist")
	GroupNotFound     = errors.New("group doesn't exist")
	SubjectsNotFound  = errors.New("subjects doesn't exist")
	ClassesNotFound   = errors.New("classes doesn't exist")
	HomeworksNotFound = errors.New("homeworks doesn't exist")
)

var (
	UserExists         = errors.New("user already exists")
	GroupExists        = errors.New("group already exists")
	ErrInvalidPassword = errors.New("invalid password provided")
)
