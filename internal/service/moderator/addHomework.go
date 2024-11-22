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
	HomeworkText   string
	DueDate        time.Time
}

// TODO: добавить проверку на то что такое занятие существует
func (s *Service) AddGroup(ctx context.Context, req AddHomework) error {
	homework := req.toHomework()

	err := s.homeworkService.Create(ctx, homework)

	if err != nil {
		return err
	}

	return nil
}

func (r *AddHomework) toHomework() entity.Homework {
	return entity.Homework{
		ClassSemNumber: r.ClassSemNumber,
		GroupID:        r.GroupID,
		SubjectID:      r.SubjectID,
		HomeworkText:   r.HomeworkText,
		DueDate:        r.DueDate,
	}
}
