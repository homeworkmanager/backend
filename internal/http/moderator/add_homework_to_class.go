package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	moderatorService "homeworktodolist/internal/service/moderator"
	"time"
)

type AddHomeworkToClassReq struct {
	ClassSemNumber int64            `json:"classSemNumber"`
	SubjectID      entity.SubjectID `json:"subjectId"`
	Category       string           `json:"category"`
	HomeworkText   string           `json:"homeworkText"`
	DueDate        time.Time        `json:"dueDate"`
}

func (h *Handler) AddHomeworkToClass() fiber.Handler {
	return func(c *fiber.Ctx) error {

		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req AddHomeworkToClassReq

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.ClassSemNumber == 0 || req.SubjectID == 0 || req.HomeworkText == "" || req.DueDate.IsZero() {
			return fiber.ErrBadRequest
		}

		category, ok := entity.CategoryToNumber[req.Category]
		if !ok {
			return fiber.ErrBadRequest
		}

		id, err := h.moderatorService.AddHomework(c.Context(), moderatorService.AddHomework{
			ClassSemNumber: &req.ClassSemNumber,
			GroupID:        creds.GroupID,
			SubjectID:      req.SubjectID,
			Category:       &category,
			HomeworkText:   req.HomeworkText,
			DueDate:        req.DueDate,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"homework_id": id,
			"data":        "Homework successfully added",
		})
	}
}
