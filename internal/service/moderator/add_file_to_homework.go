package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
	"mime/multipart"
)

type AddFileReq struct {
	FileHeader *multipart.FileHeader
	HomeworkID entity.HomeworkID
	GroupID    entity.GroupID
}

func (s *Service) AddFileToHomework(ctx context.Context, req AddFileReq) (entity.FileID, error) {
	id, err := s.homeworkFileService.AddFileToHomework(ctx, req.FileHeader, req.HomeworkID, req.GroupID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
