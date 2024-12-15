package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	moderatorService "homeworktodolist/internal/service/moderator"
	"time"
)

type AddHomeworkToDateReq struct {
	GroupID      entity.GroupID   `json:"groupId"`
	SubjectID    entity.SubjectID `json:"subjectId"`
	HomeworkText string           `json:"homeworkText"`
	DueDate      time.Time        `json:"dueDate"`
}

func (h *Handler) AddHomeworkToDate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AddHomeworkToDateReq

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.GroupID == 0 || req.SubjectID == 0 || req.HomeworkText == "" || req.DueDate.IsZero() {
			return fiber.ErrBadRequest
		}

		err := h.moderatorService.AddHomework(c.Context(), req.toAdd())
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": "Homework successfully added",
		})
	}
}
func (r *AddHomeworkToDateReq) toAdd() moderatorService.AddHomework {
	return moderatorService.AddHomework{
		GroupID:      r.GroupID,
		SubjectID:    r.SubjectID,
		HomeworkText: r.HomeworkText,
		DueDate:      r.DueDate,
	}
}
