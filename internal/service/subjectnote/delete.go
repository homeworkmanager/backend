package subjectnote

import (
	"context"
	"homeworktodolist/internal/entity"
)

func (s *Service) Delete(ctx context.Context, noteID entity.NoteID) error {
	err := s.subjectNoteRepo.Delete(ctx, noteID)
	if err != nil {
		return err
	}
	return nil
}
