package homework_files

import (
	"context"
	"mime/multipart"
)

func (s *Service) Create(ctx context.Context, file *multipart.File) error {
	keyName := "4"
	err := s.homeworkFilesS3.Upload(ctx, file, keyName)
	if err != nil {
		return err
	}
	return nil
}
