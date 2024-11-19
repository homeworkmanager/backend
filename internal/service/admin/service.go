package admin

type Service struct {
	groupService   GroupService
	classService   ClassService
	subjectService SubjectService
}

func NewAdminService(groupService GroupService, classService ClassService, subjectService SubjectService) *Service {
	return &Service{
		groupService:   groupService,
		classService:   classService,
		subjectService: subjectService,
	}
}
