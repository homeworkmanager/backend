package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) RegenerateGroupRegisterKey(ctx context.Context, groupID entity.GroupID) error {
	err := s.groupService.RegenerateGroupRegisterKey(ctx, groupID)
	if err != nil {
		return err
	}
	return nil
}
