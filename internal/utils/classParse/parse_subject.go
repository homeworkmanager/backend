package classParse

import (
	"homeworktodolist/internal/entity"
	"strings"
)

func ParseSubject(classes []entity.UpdateClass) []string {
	uniqueMap := make(map[string]struct{})

	var res []string

	for _, class := range classes {
		subjectName := strings.Join(strings.Split(class.Summary, " ")[1:], " ")
		if _, exists := uniqueMap[subjectName]; !exists {
			uniqueMap[subjectName] = struct{}{}
			res = append(res, subjectName)
		}
	}

	return res
}
