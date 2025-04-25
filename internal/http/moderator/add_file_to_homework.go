package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	moderatorService "homeworktodolist/internal/service/moderator"
	"homeworktodolist/internal/utils"
	"strconv"
)

func (h *Handler) AddFileToHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {
		creds, ok := c.Locals(entity.Claims).(entity.UserCreds)
		if !ok {
			return fiber.ErrUnauthorized
		}

		homeworkID, err := strconv.Atoi(c.Params("homeworkID"))
		if err != nil {
			return fiber.ErrBadRequest
		}

		fileHeader, err := c.FormFile("file")
		if err != nil {
			return fiber.ErrBadRequest
		}

		if err = utils.CheckFile(fileHeader); err != nil {
			return fiber.ErrBadRequest
		}

		id, err := h.moderatorService.AddFileToHomework(c.Context(), moderatorService.AddFileReq{
			FileHeader: fileHeader,
			HomeworkID: entity.HomeworkID(homeworkID),
			GroupID:    creds.GroupID,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"file_id": id,
			"data":    "File successfully added",
		})

	}
}
