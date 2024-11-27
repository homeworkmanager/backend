package homework

type Service struct {
	homeworkRepo HomeworkRepo
}

func NewHomeworkService(homeworkRepo HomeworkRepo) *Service {
	return &Service{
		homeworkRepo: homeworkRepo,
	}
}
