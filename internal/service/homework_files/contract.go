package homework_files

import (
	"context"
	"homeworktodolist/internal/entity"
	"mime/multipart"
)

type HomeworkFilesS3 interface {
	Upload(ctx context.Context, file *multipart.FileHeader, key string) error
	Delete(ctx context.Context, key string) error
}
type HomeworkFilesRepo interface {
	Create(ctx context.Context, HomeworkFile entity.HomeworkFile) (entity.FileID, error)
	GetByFileID(ctx context.Context, fileID entity.FileID) (entity.HomeworkFile, error)
	Delete(ctx context.Context, fileID entity.FileID) error
}
