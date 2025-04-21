package homework_files

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (r *Repo) Create(ctx context.Context, HomeworkFile entity.HomeworkFile) error {
	q := "INSERT INTO homeworksfiles (homework_id, file_name, file_url) values ($1, $2, $3)"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, HomeworkFile.HomeworkID, HomeworkFile.FileName, HomeworkFile.FileURL)
	if err != nil {
		return err
	}
	return nil
}
