package homework_files

import (
	"context"
	"mime/multipart"
)

type HomeworkFilesS3 interface {
	Upload(ctx context.Context, file *multipart.File, keyName string) error
}
