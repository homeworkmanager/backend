package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) DeleteNote(ctx context.Context, id entity.NoteID) error {
	err := s.subjectNoteService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
