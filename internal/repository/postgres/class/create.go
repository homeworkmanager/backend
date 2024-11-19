package class

import (
	"context"
	"fmt"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, classes []entity.Class) error {

	//TODO: заюзать здесь sql билдер
	q := "INSERT INTO classes(group_id, subject_id, start_time, end_time, summary, class_sem_number, location) VALUES "

	var args []interface{}

	for i, class := range classes {
		if i > 0 {
			q += ", "
		}
		q += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*7+1, i*7+2, i*7+3, i*7+4, i*7+5, i*7+6, i*7+7)
		args = append(args, class.GroupID, class.SubjectID, class.StartTime, class.EndTime, class.Summary, class.SemClassNumber, class.Location)
	}

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil

}
