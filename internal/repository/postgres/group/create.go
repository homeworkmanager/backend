package group

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, group entity.Group) (entity.GroupID, error) {
	var id entity.GroupID

	q := "INSERT INTO groups(group_name, course, ical_link) VALUES ($1, $2, $3) RETURNING group_id"

	t := r.manager.GetTxOrDefault(ctx)

	err := t.QueryRowContext(ctx, q, group.Name, group.Course, group.IcalLink).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err

}
