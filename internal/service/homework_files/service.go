package homework_files

import "homeworktodolist/internal/client/http/s3"

type Service struct {
	homeworkFilesS3 HomeworkFilesS3
}

func NewHomeworkFilesService(client *s3.S3Client) *Service {
	return &Service{
		client,
	}
}
