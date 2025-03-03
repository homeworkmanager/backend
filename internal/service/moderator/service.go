package moderator

type Service struct {
	homeworkService    HomeworkService
	subjectNoteService SubjectNoteService
	groupService       GroupService
}

func NewModeratorService(homeworkService HomeworkService, subjectNoteService SubjectNoteService, groupService GroupService) *Service {
	return &Service{
		homeworkService:    homeworkService,
		subjectNoteService: subjectNoteService,
		groupService:       groupService,
	}
}
