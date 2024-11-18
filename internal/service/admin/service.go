package admin

type Service struct {
	groupService GroupService
	classService ClassService
}

func NewAdminService(groupService GroupService, classService ClassService) *Service {
	return &Service{
		groupService: groupService,
		classService: classService,
	}
}
