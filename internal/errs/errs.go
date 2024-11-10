package errs

import "errors"

var (
	UserNotFound = errors.New("user doesn't exist")
	UserExists   = errors.New("user already exists")
)
