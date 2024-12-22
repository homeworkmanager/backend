package schedule

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	scheduleService "homeworktodolist/internal/service/schedule"
	"time"
)

type GetReq struct {
	FromTime  string `query:"from_time"`
	DaysCount int    `query:"days_count"`
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

		if req.FromTime == "" || req.DaysCount == 0 {
			return fiber.ErrBadRequest
		}

		parsedFromTime, err := time.Parse("02.01.2006", req.FromTime)
		if err != nil {
			return fiber.ErrBadRequest
		}

		days, err := h.scheduleService.GetAllByGroupAndTime(c.Context(), scheduleService.GetSchedule{
			GroupId:   creds.GroupID,
			FromTime:  parsedFromTime,
			DaysCount: req.DaysCount,
		})
		if err != nil {
			return err
		}

		daysMap := make(map[string]day)
		for _, d := range days {
			key := d.Date.Format("2006-01-02")
			daysMap[key] = day{
				OutputClass:         toOutputClass(d.OutputClass),
				IndependentHomework: toHomework(d.IndependentHomework),
			}
		}
		return c.JSON(daysMap)
	}
}
