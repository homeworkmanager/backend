package entity

const (
	RoleUser Role = iota + 1
	RoleGroupModerator
	RoleGlobalAdmin
)

const (
	SessionKey = "session_key"
	Claims     = "claims"
)

const (
	CategoryLection      = "ЛК"
	CategoryPractice     = "ПР"
	CategoryLab          = "ЛАБ"
	CategoryCredit       = "Зачет"
	CategoryConsultation = "Консультация"
	CategoryExam         = "Экзамен"
)

var CategoryToNumber = map[string]ClassCategory{
	CategoryLection:      1,
	CategoryPractice:     2,
	CategoryLab:          3,
	CategoryCredit:       4,
	CategoryConsultation: 5,
	CategoryExam:         6,
}

var NumberToCategory = map[ClassCategory]string{
	1: CategoryLection,
	2: CategoryPractice,
	3: CategoryLab,
	4: CategoryCredit,
	5: CategoryConsultation,
	6: CategoryExam,
}
