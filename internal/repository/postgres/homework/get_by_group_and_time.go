package homework

import (
	"context"
	"time"

	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

func (r *Repo) GetByGroupAndTime(ctx context.Context, userID entity.UserID, groupID entity.GroupID, fromTime time.Time, toTime time.Time) ([]entity.Homework, error) {
	q := "WITH tmp AS(SELECT homeworks.*, COALESCE(homeworkstatuses.id IS NOT NULL AND homeworkstatuses.id != 0, TRUE) AS Status " +
		"FROM homeworks LEFT OUTER JOIN  homeworkstatuses " +
		"ON homeworks.homework_id = homeworkstatuses.homework_id " +
		"WHERE homeworks.group_id = $1 AND homeworks.due_date >= $2 " +
		"AND homeworks.due_date <= $3 " +
		"AND (homeworkstatuses.user_id = $4 OR homeworkstatuses.user_id IS NULL ))" +
		"SELECT tmp.*, subject_name FROM tmp JOIN subjects ON tmp.subject_id = subjects.subject_id"

	var homeworks []homework

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &homeworks, q, groupID, fromTime, toTime, userID)
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
