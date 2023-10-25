package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (t *Transport) Sign(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(response)
}
