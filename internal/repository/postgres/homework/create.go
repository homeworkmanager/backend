package homework

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, homework entity.Homework) error {

	q := "INSERT INTO homeworks(class_sem_number, group_id, subject_id, homework_text, category, due_date) VALUES ($1, $2, $3, $4, $5, $6)"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, homework.SemClassNumber, homework.GroupID, homework.SubjectID, homework.HomeworkText, homework.Category, homework.DueDate)

	if err != nil {
		return err
	}
	return nil

}
