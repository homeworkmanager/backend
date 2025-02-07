package admin

import "homeworktodolist/internal/tx_manager"

type Service struct {
	groupService       GroupService
	classService       ClassService
	subjectService     SubjectService
	homeworkService    HomeworkService
	userService        UserService
	subjectNoteService SubjectNoteService
	manager            *tx_manager.TxManager
}

func NewAdminService(groupService GroupService, classService ClassService, subjectService SubjectService, homeworkService HomeworkService, userService UserService, subjectNoteService SubjectNoteService, manager *tx_manager.TxManager) *Service {
	return &Service{
		groupService:       groupService,
		classService:       classService,
		subjectService:     subjectService,
		homeworkService:    homeworkService,
		manager:            manager,
		userService:        userService,
		subjectNoteService: subjectNoteService,
	}
}
