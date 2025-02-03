package moderator

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) UpdateNote(ctx context.Context, noteID entity.NoteID, noteText string) error {
	err := s.subjectNoteService.Update(ctx, noteID, noteText)
	if err != nil {
		return err
	}
	return nil
}
