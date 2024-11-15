package group

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, group entity.Group) error {

	t := r.manager.GetTxOrDefault(ctx)
	q := "INSERT INTO groups(group_name, course, ical_link) VALUES ($1, $2, $3);"

	_, err := t.ExecContext(ctx, q, group.Name, group.Course, group.IcalLink)
	if err != nil {
		return err
	}

	return nil

}
