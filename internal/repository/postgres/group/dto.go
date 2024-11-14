package group

import (
	"homeworktodolist/internal/entity"
)

type group struct {
	GroupID entity.GroupID `db:"group_id"`
	Name    string         `db:"group_name"`
	Course  int8           `db:"course"`
}

func (g group) toGroup() entity.Group {
	return entity.Group{
		GroupID: g.GroupID,
		Name:    g.Name,
		Course:  g.Course,
	}
}
