package user

import (
	"context"

	"homeworktodolist/internal/entity"
)

func (s Service) GetUserFull(ctx context.Context, id entity.UserID) (entity.UserFullInfo, error) {
	user, err := s.userRepo.GetFullById(ctx, id)
	if err != nil {
		return entity.UserFullInfo{}, err
	}
	return user, nil
}
