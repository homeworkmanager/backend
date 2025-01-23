package user

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) UpdateRole(ctx context.Context, userID entity.UserID, role entity.Role) error {
	err := s.userRepo.UpdateRole(ctx, userID, role)
	if err != nil {
		return err
	}
	return nil
}
