package classParse

import (
	"homeworktodolist/internal/entity"
)

func ParseSubject(classes []entity.UpdateClass) []string {
	uniqueMap := make(map[string]struct{})

	var res []string

	for _, class := range classes {

		subject := string([]rune(class.Summary)[3:])
		if _, exists := uniqueMap[subject]; !exists {
			uniqueMap[subject] = struct{}{}
			res = append(res, subject)
		}
	}

	return res
}
