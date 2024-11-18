package admin

import (
	"context"
)

func (s *Service) UpdateClasses(ctx context.Context) error {
	err := s.classService.Update(ctx)
	return err
}
