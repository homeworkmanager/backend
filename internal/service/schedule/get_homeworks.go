package schedule

import (
	"context"
	"time"

	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
	"homeworktodolist/internal/utils"
)

type GetHomework struct {
	UserID    entity.UserID
	GroupID   entity.GroupID
	FromTime  time.Time
	DaysCount int
}

func (s *Service) GetHomeworks(ctx context.Context, req GetHomework) ([]entity.HomeworkDay, error) {
	toTime := req.FromTime.Add(time.Duration(req.DaysCount*24) * time.Hour)

	homeworks, err := s.homeworkService.GetByGroupAndTime(ctx, req.UserID, req.GroupID, req.FromTime, toTime)
	if err != nil {
		if err != errs.HomeworksNotFound {
			return nil, err
		}
	}

	days := make([]entity.HomeworkDay, req.DaysCount)

	for i := 0; i < req.DaysCount; i++ {
		days[i] = entity.HomeworkDay{
			Date:     req.FromTime.Add(time.Duration(i) * 24 * time.Hour).Local(),
			Homework: []entity.Homework{},
		}
	}

	homeworkMap := make(map[time.Time][]entity.Homework)
	for _, homework := range homeworks {
		date := time.Date(
			homework.DueDate.Year(), homework.DueDate.Month(), homework.DueDate.Day(),
			0, 0, 0, 0,
			time.Local)
		homeworkMap[date] = append(homeworkMap[date], homework)
	}

	for i, day := range days {
		dayHomeworks := homeworkMap[day.Date]
		for _, homework := range dayHomeworks {
			if utils.DeRef[int64](homework.SemClassNumber) == 0 {
				day.Homework = append(day.Homework, homework)
			}
		}
		days[i] = day
	}
	return days, nil
}
