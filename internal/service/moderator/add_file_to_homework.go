package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
	"mime/multipart"
)

type AddFileReq struct {
	FilesHeader []*multipart.FileHeader
	HomeworkID  entity.HomeworkID
	GroupID     entity.GroupID
}

func (s *Service) AddFileToHomework(ctx context.Context, req AddFileReq) (map[string]entity.FileID, map[entity.FileID]string, map[string]error, error) {
	filesIdMap := make(map[string]entity.FileID)
	fileURLMap := make(map[entity.FileID]string)
	filesErrMap := make(map[string]error)

	//TODO: добавить параллельность
	//TODO: Написать в доке о том, что файл может не добавиться и тогда его имени просто не будет в мапке которую я возвращаю
	for _, file := range req.FilesHeader {
		fileID, fileURL, err := s.homeworkFileService.AddFileToHomework(ctx, file, req.HomeworkID, req.GroupID)
		if err != nil {
			filesErrMap[file.Filename] = err
			continue
		}
		filesIdMap[file.Filename] = fileID
		fileURLMap[fileID] = fileURL
	}

	return filesIdMap, fileURLMap, filesErrMap, nil
}
