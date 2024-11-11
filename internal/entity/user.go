package entity

import "time"

type User struct {
	UserID    UserID
	Name      string
	Surname   *string
	Email     string
	Password  string
	Role      Role
	GroupId   int64
	CreatedAt time.Time
}

type UserCreds struct {
	UserID  UserID
	Role    Role
	GroupId int64
}
