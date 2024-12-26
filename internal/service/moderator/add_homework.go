package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
	"time"
)

type AddHomework struct {
	ClassSemNumber *int64
	GroupID        entity.GroupID
	SubjectID      entity.SubjectID
	Category       *entity.ClassCategory
	HomeworkText   string
	DueDate        time.Time
}

// TODO: добавить проверку на то что такое занятие существует
func (s *Service) AddHomework(ctx context.Context, req AddHomework) error {
	homework := req.toHomework()

	err := s.homeworkService.Create(ctx, homework)

	if err != nil {
		return err
	}

	return nil
}

func (r *AddHomework) toHomework() entity.Homework {
	return entity.Homework{
		SemClassNumber: r.ClassSemNumber,
		GroupID:        r.GroupID,
		SubjectID:      r.SubjectID,
		Category:       r.Category,
		HomeworkText:   r.HomeworkText,
		DueDate:        r.DueDate,
	}
}
