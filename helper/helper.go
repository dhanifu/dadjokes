package helper

import (
	"github.com/dhanifu/dadjokes/models"
	"github.com/gofiber/fiber/v2"
)

// ResponseSuccess simplify fiber success function
func ResponseSuccess(c *fiber.Ctx, data interface{}, messages ...string) error {
	message := "Data berhasil didapatkan"
	if len(messages) > 0 {
		message = messages[0]
	}

	return c.JSON(models.Response{
		Success:    true,
		Data:       data,
		Message:    message,
	})
}