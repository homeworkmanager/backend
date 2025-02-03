package moderator

type Service struct {
	homeworkService    HomeworkService
	subjectNoteService SubjectNoteService
}

func NewModeratorService(homeworkService HomeworkService, subjectNoteService SubjectNoteService) *Service {
	return &Service{
		homeworkService:    homeworkService,
		subjectNoteService: subjectNoteService,
	}
}
