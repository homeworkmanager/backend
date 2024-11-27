package schedule

import (
	"context"
	"database/sql"
	"errors"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
)

func (r *Repo) GetByGroupId(ctx context.Context, groupId entity.GroupID) ([]entity.Schedule, error) {

	q := "SELECT classes.class_id, COALESCE(classes.group_id, homeworks.group_id) AS group_id, COALESCE(classes.subject_id, homeworks.subject_id) AS subject_id, classes.start_time, classes.end_time, classes.summary, COALESCE(classes.class_sem_number, homeworks.class_sem_number) AS class_sem_number , classes.location, homework_id, homework_text, due_date FROM classes\n" +
		"FULL OUTER JOIN homeworks on classes.group_id = homeworks.group_id and classes.subject_id = homeworks.subject_id\n    and classes.class_sem_number = homeworks.class_sem_number\n" +
		"WHERE  (classes.group_id = $1 OR homeworks.group_id = $1)\n" +
		"AND (classes.start_time >= '2024-09-02' OR classes.start_time IS NULL)\n" +
		"AND (classes.end_time < '2024-09-09' OR classes.end_time IS NULL)\n"

	var schedule []schedule

	t := r.manager.GetTxOrDefault(ctx)

	err := t.SelectContext(ctx, &schedule, q, groupId)
	if errors.Is(err, sql.ErrNoRows) {
		return []entity.Schedule{}, errs.ClassesNotFound
	}
	if err != nil {
		return []entity.Schedule{}, err
	}

	return toSchedules(schedule), nil

}
