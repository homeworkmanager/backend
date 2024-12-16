package schedule

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/service/schedule"
	"time"
)

type GetReq struct {
	FromTime  time.Time `json:"from_time"`
	CountDays int       `json:"count_days"`
}

func (h *Handler) GetSchedule() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		var req GetReq
		if err := c.QueryParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if req.FromTime.IsZero() || req.CountDays == 0 {
			return fiber.ErrBadRequest
		}

		days, err := h.scheduleService.GetAllByGroupAndTime(c.Context(), schedule.GetSchedule{
			GroupId:   creds.GroupID,
			FromTime:  req.FromTime,
			CountDays: req.CountDays,
		})
		if err != nil {
			return err
		}

		daysMap := make(map[string]day)
		for _, d := range days {
			key := d.Date.Format("2006-01-02") // Форматируем дату в строку
			daysMap[key] = day{
				OutPutClass:         toOutputClass(d.OutPutClass),
				IndependentHomework: toHomework(d.IndependentHomework),
			}
		}
		return c.JSON(daysMap)
	}
}
