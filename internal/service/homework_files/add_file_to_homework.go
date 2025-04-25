package homework_files

import (
	"context"
	"fmt"
	"homeworktodolist/internal/entity"
	"mime/multipart"
)

func (s *Service) AddFileToHomework(ctx context.Context, fileHeader *multipart.FileHeader, homeworkID entity.HomeworkID, groupID entity.GroupID) (entity.FileID, error) {
	key := fmt.Sprintf("homeworks/%d/%s", homeworkID, fileHeader.Filename)
	err := s.homeworkFilesS3.Upload(ctx, fileHeader, key)
	if err != nil {
		return 0, err
	}
	fileName := fileHeader.Filename
	fileUrl := "https://global.s3.cloud.ru/unihelper/" + key + "/" + fileName
	homeworkFile := entity.HomeworkFile{
		HomeworkID: homeworkID,
		FileName:   fileName,
		GroupID:    groupID,
		FileURL:    fileUrl,
		Key:        key,
	}
	id, err := s.homeworkFilesRepo.Create(ctx, homeworkFile)
	if err != nil {
		return 0, err
	}
	return id, err
}
