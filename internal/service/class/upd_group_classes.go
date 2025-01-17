package class

import (
	"context"
	"strings"

	"homeworktodolist/internal/entity"
)

type classCount struct {
	entity.ClassCategory
	entity.SubjectID
}

func (s *Service) UpdGroupClasses(ctx context.Context, group entity.Group, classes []entity.Class) error {
	var err error

	classCounter := make(map[classCount]int64)

	for i := range classes {

		subject, err := s.subjectService.GetBySubjectNameAndGroup(ctx, strings.Join(strings.Split(classes[i].Summary, " ")[1:], " "), group.GroupID)
		if err != nil {
			return err
		}

		classC := classCount{
			ClassCategory: classes[i].Category,
			SubjectID:     subject.SubjectId,
		}

		classCounter[classC]++

		classes[i].SubjectID = subject.SubjectId
		classes[i].GroupID = group.GroupID
		classes[i].SemClassNumber = classCounter[classC]
	}

	err = s.classRepo.Create(ctx, classes)

	if err != nil {
		return err
	}

	return nil

}
