package schedule

import (
	"context"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
	"homeworktodolist/internal/utils"
	"time"
)

type GetSchedule struct {
	GroupId   entity.GroupID
	FromTime  time.Time
	CountDays int
}

func (s *Service) GetAllByGroupAndTime(ctx context.Context, req GetSchedule) ([]entity.Day, error) {
	toTime := req.FromTime.Add(time.Duration(req.CountDays*24) * time.Hour)

	classes, err := s.classService.GetByGroupAndTime(ctx, req.GroupId, req.FromTime, toTime)
	if err != nil {
		if err != errs.ClassesNotFound {
			return nil, err
		}
	}

	homeworks, err := s.homeworkService.GetByGroupAndTime(ctx, req.GroupId, req.FromTime, toTime)
	if err != nil {
		if err != errs.HomeworksNotFound {
			return nil, err
		}
	}

	days := make([]entity.Day, req.CountDays)

	for i := 0; i < req.CountDays; i++ {
		days[i] = entity.Day{
			Date:                req.FromTime.Add(time.Duration(i) * 24 * time.Hour),
			OutPutClass:         []entity.OutputClass{},
			IndependentHomework: []entity.Homework{},
		}
	}

	classMap := make(map[time.Time][]entity.Class)
	for _, class := range classes {
		date := class.StartTime.Truncate(24 * time.Hour)
		classMap[date] = append(classMap[date], class)
	}

	homeworkMap := make(map[time.Time][]entity.Homework)
	for _, homework := range homeworks {
		date := homework.DueDate.Truncate(24 * time.Hour)
		homeworkMap[date] = append(homeworkMap[date], homework)
	}

	for i, day := range days {
		dayClasses := classMap[day.Date]
		dayHomeworks := homeworkMap[day.Date]

		for _, class := range dayClasses {
			outputClass := entity.OutputClass{
				Class:    class,
				Homework: []entity.Homework{},
			}

			for _, homework := range dayHomeworks {
				if homework.SubjectID == class.SubjectID && utils.DeRef[int64](homework.ClassSemNumber) == class.SemClassNumber {
					outputClass.Homework = append(outputClass.Homework, homework)
				}
			}

			day.OutPutClass = append(day.OutPutClass, outputClass)
		}

		for _, homework := range dayHomeworks {
			if utils.DeRef[int64](homework.ClassSemNumber) == 0 {
				day.IndependentHomework = append(day.IndependentHomework, homework)
			}
		}

		days[i] = day

	}
	return days, nil
}
