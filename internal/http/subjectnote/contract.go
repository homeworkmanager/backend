package subjectnote

import (
	"context"
	"homeworktodolist/internal/entity"
)

type SubjectNoteService interface {
	GetByGroup(ctx context.Context, groupID entity.GroupID) ([]entity.SubjectWithNotes, error)
}
