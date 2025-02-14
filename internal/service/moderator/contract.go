package moderator

import (
	"context"

	"homeworktodolist/internal/entity"
)

type HomeworkService interface {
	Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error)
	Delete(ctx context.Context, id entity.HomeworkID) error
	Update(ctx context.Context, homeworkId entity.HomeworkID, homeworkText string) error
}
type SubjectNoteService interface {
	AddNote(ctx context.Context, note entity.SubjectNote) (entity.NoteID, error)
	Delete(ctx context.Context, noteID entity.NoteID) error
	Update(ctx context.Context, noteID entity.NoteID, noteText string) error
}

type GroupService interface {
	RegenerateGroupRegisterKey(ctx context.Context, groupID entity.GroupID) (string, error)
	GetByID(ctx context.Context, groupID entity.GroupID) (entity.Group, error)
}
