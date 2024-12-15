package schedule

type Service struct {
	classService    ClassService
	homeworkService HomeworkService
}

func NewScheduleService(classService ClassService, homeworkService HomeworkService) *Service {
	return &Service{
		classService:    classService,
		homeworkService: homeworkService,
	}
}
