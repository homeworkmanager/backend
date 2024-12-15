package homework

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
	"time"
)

func (r *Repo) GetByGroupAndTime(ctx context.Context, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error) {
	q := "SELECT * FROM homeworks WHERE group_id = $1 AND due_date >= $2 AND due_date <= $3"

	var homeworks []homework

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &homeworks, q, groupID, fromTime, toTime)
	if len(homeworks) == 0 {
		return []entity.Homework{}, errs.HomeworksNotFound
	}
	if err != nil {
		return []entity.Homework{}, err
	}

	return toHomeworks(homeworks), nil
}

func toHomeworks(homeworks []homework) []entity.Homework {
	res := make([]entity.Homework, len(homeworks))
	for i, g := range homeworks {
		res[i] = g.toHomework()
	}
	return res
}
