package group

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, group entity.Group) error {

	t := r.manager.GetTxOrDefault(ctx)
	q := "INSERT INTO groups(group_name, course) VALUES ($1, $2);"

	_, err := t.ExecContext(ctx, q, group.Name, group.Course)
	if err != nil {
		return err
	}

	return nil

}
