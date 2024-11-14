package admin

type Service struct {
	groupService GroupService
}

func NewAdminService(groupService GroupService) *Service {
	return &Service{
		groupService: groupService,
	}
}
