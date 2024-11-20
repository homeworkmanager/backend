package admin

import "homeworktodolist/internal/tx_manager"

type Service struct {
	groupService   GroupService
	classService   ClassService
	subjectService SubjectService
	manager        *tx_manager.TxManager
}

func NewAdminService(groupService GroupService, classService ClassService, subjectService SubjectService, manager *tx_manager.TxManager) *Service {
	return &Service{
		groupService:   groupService,
		classService:   classService,
		subjectService: subjectService,
		manager:        manager,
	}
}
