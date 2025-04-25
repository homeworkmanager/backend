package homework_files

import (
	"homeworktodolist/internal/client/http/s3"
	"homeworktodolist/internal/tx_manager"
)

type Service struct {
	homeworkFilesS3   HomeworkFilesS3
	homeworkFilesRepo HomeworkFilesRepo
	manager           *tx_manager.TxManager
}

func NewHomeworkFilesService(client *s3.S3Client, homeworkFilesRepo HomeworkFilesRepo, manager *tx_manager.TxManager) *Service {
	return &Service{
		client,
		homeworkFilesRepo,
		manager,
	}
}
