package moderator

import (
	"context"
	"mime/multipart"
	"time"

	"homeworktodolist/internal/entity"
)

type AddHomework struct {
	ClassSemNumber *int64
	GroupID        entity.GroupID
	SubjectID      entity.SubjectID
	Category       *entity.ClassCategory
	HomeworkText   string
	FilesHeader    []*multipart.FileHeader
	DueDate        time.Time
}

// TODO: добавить проверку на то что такое занятие существует
func (s *Service) AddHomework(ctx context.Context, req AddHomework) (entity.HomeworkID, map[string]entity.FileID, error) {
	homework := req.toHomework()
	id, err := s.homeworkService.Create(ctx, homework)
	if err != nil {
		return 0, nil, err
	}

	filesIdMap := make(map[string]entity.FileID)

	//TODO: добавить параллельность
	//TODO: Написать в доке о том, что файл может не добавиться и тогда его имени просто не будет в мапке которую я возвращаю
	for _, file := range req.FilesHeader {
		fileID, err := s.homeworkFileService.AddFileToHomework(ctx, file, id, req.GroupID)
		if err != nil {
			continue
		}
		filesIdMap[file.Filename] = fileID
	}

	return id, filesIdMap, err
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
