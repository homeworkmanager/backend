package user

import (
	"homeworktodolist/internal/entity"
	"time"
)

type user struct {
	UserID    entity.UserID  `db:"user_id"`
	Name      string         `db:"name"`
	Surname   *string        `db:"surname"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Role      entity.Role    `db:"role"`
	GroupID   entity.GroupID `db:"group_id"`
	CreatedAt time.Time      `db:"created_at"`
}

func (u user) toUser() entity.User {
	return entity.User{
		UserID:    u.UserID,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		GroupID:   u.GroupID,
		CreatedAt: u.CreatedAt,
	}
}
