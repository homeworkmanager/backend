package subject

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
)

type SubjectResp struct {
	SubjectID   entity.SubjectID `json:"subject_id"`
	SubjectName string           `json:"subject_name"`
}

func (h *Handler) GetSubjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}
		s, err := h.subjectService.GetByGroupId(c.Context(), creds.GroupID)
		if err != nil {
			return err
		}
		subjects := toSubjectResp(s)
		return c.JSON(subjects)
	}
}

func toSubjectResp(g []entity.Subject) []SubjectResp {
	res := make([]SubjectResp, len(g))
	for i := range g {
		res[i] = SubjectResp{
			SubjectID:   g[i].SubjectId,
			SubjectName: g[i].SubjectName,
		}
	}
	return res
}
