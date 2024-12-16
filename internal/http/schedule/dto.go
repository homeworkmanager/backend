package schedule

import (
	"homeworktodolist/internal/entity"
	"time"
)

type homework struct {
	HomeworkID     entity.HomeworkID `json:"homeworkID"`
	ClassSemNumber *int64            `json:"classSemNumber"`
	SubjectID      entity.SubjectID  `json:"subjectId"`
	HomeworkText   string            `json:"homeworkText"`
	DueDate        time.Time         `json:"dueDate"`
}

type class struct {
	SubjectID      entity.SubjectID `json:"subjectId"`
	StartTime      time.Time        `json:"startTime"`
	EndTime        time.Time        `json:"endTime"`
	Summary        string           `json:"summary"`
	SemClassNumber int64            `json:"semClassNumber"`
	Location       string           `json:"location"`
}

type outputClass struct {
	Class    class      `json:"class"`
	Homework []homework `json:"homework"`
}

type day struct {
	OutputClass         []outputClass `json:"outputClasses"`
	IndependentHomework []homework    `json:"independentHomeworks"`
}

func toClass(c entity.Class) class {
	return class{
		SubjectID:      c.SubjectID,
		StartTime:      c.StartTime,
		EndTime:        c.EndTime,
		Summary:        c.Summary,
		SemClassNumber: c.SemClassNumber,
		Location:       c.Location,
	}
}
func toHomework(c []entity.Homework) []homework {
	result := make([]homework, len(c))
	for i, h := range c {
		result[i] = homework{
			HomeworkID:     h.HomeworkID,
			ClassSemNumber: h.SemClassNumber,
			SubjectID:      h.SubjectID,
			HomeworkText:   h.HomeworkText,
			DueDate:        h.DueDate,
		}
	}
	return result
}

func toOutputClass(c []entity.OutputClass) []outputClass {
	result := make([]outputClass, len(c))
	for i, h := range c {
		result[i] = outputClass{
			Class:    toClass(h.Class),
			Homework: toHomework(h.Homework),
		}
	}
	return result
}
