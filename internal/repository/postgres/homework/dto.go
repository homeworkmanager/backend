package homework

import (
	"homeworktodolist/internal/entity"
	"time"
)

type homework struct {
	HomeworkID     entity.HomeworkID     `db:"homework_id"`
	ClassSemNumber *int64                `db:"class_sem_number"`
	GroupID        entity.GroupID        `db:"group_id"`
	SubjectID      entity.SubjectID      `db:"subject_id"`
	Category       *entity.ClassCategory `db:"category"`
	HomeworkText   string                `db:"homework_text"`
	DueDate        time.Time             `db:"due_date"`
	CreatedAt      time.Time             `db:"created_at"`
}

func (h homework) toHomework() entity.Homework {
	return entity.Homework{
		HomeworkID:     h.HomeworkID,
		SemClassNumber: h.ClassSemNumber,
		GroupID:        h.GroupID,
		SubjectID:      h.SubjectID,
		Category:       h.Category,
		HomeworkText:   h.HomeworkText,
		DueDate:        h.DueDate,
		CreatedAt:      h.CreatedAt,
	}
}
