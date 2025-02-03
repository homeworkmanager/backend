package user

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) GetAllUsersFull(ctx context.Context) ([]entity.UserFullInfo, error) {
	users, err := s.userRepo.GetAllFull(ctx)
	if err != nil {
		return []entity.UserFullInfo{}, err
	}
	return users, nil
}
