package classParse

import "homeworktodolist/internal/service/class"

func ParseSubject(classes []class.UpdateClass) []string {
	uniqueMap := make(map[string]struct{})

	var res []string

	for _, class := range classes {
		subject := class.Summary[3:]
		if _, exists := uniqueMap[subject]; !exists {
			uniqueMap[subject] = struct{}{}
			res = append(res, subject)
		}
	}

	return res
}
