package moderator

type Service struct {
	homeworkService HomeworkService
}

func NewModeratorService(homeworkService HomeworkService) *Service {
	return &Service{
		homeworkService: homeworkService,
	}
}
