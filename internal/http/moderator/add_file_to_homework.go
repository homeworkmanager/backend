package moderator

import (
	"github.com/gofiber/fiber/v2"
	"homeworktodolist/internal/entity"
	"homeworktodolist/internal/errs"
	"path/filepath"
	"strings"
)

func (h *Handler) AddFileToHomework() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return fiber.ErrBadRequest
		}

		if fileHeader.Size > 20*1024*1024 {
			return errs.FileTooLarge
		}
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !entity.PermittedExt[ext] {
			return errs.InvalidFileType
		}
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		err = h.homeworkFilesService.Create(c.Context(), &file)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"homework_id": "",
			"data":        "File successfully added",
		})

	}
}
