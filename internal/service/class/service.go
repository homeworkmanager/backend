package class

type Service struct {
	groupService   GroupService
	classRepo      ClassRepo
	subjectService SubjectService
}

func NewClassService(groupService GroupService, classRepo ClassRepo, subjectService SubjectService) *Service {
	return &Service{
		groupService:   groupService,
		classRepo:      classRepo,
		subjectService: subjectService,
	}
}
